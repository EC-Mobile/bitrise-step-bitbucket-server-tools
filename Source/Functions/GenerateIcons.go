package function

import (
	env "bitbucket-tools/Source/Environment"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"regexp"

	"github.com/golang/freetype"
	"golang.org/x/image/draw"
	"golang.org/x/image/font/gofont/goregular"
)

func PerformGenerateIcons() {
	fmt.Println()
	fmt.Println("Fetching Base Files....")
	fmt.Println("........................................")
	fileFilterRegex := "(.png|.jpg)$"
	baseIconsFile := getFiles(env.GI_BASE_ICONS_SET, &fileFilterRegex)
	fmt.Println("Loaded icons are:")
	for index, iconFile := range baseIconsFile {
		fmt.Printf("%d: %s\n", index, iconFile.Name())
	}

	fmt.Println()
	fmt.Println("Generating icon overlay....")
	fmt.Println("........................................")
	iconOverlay := createOverlay(
		Overlay{
			Resolution:     env.GI_ICON_OVERLAY_RESOLUTION,
			Scale:          env.GI_ICON_OVERLAY_SCALE,
			PrimaryColor:   env.GI_ICON_PRIMARY_COLOR,
			SecondaryColor: env.GI_ICON_SECONDARY_COLOR,
		},
		Label{
			Size:  float64(env.GI_ICON_OVERLAY_FONT_SIZE),
			Color: env.GI_ICON_LABEL_COLOR,
		},
		AppInfo{
			VersionNumber: env.GI_APP_VERSION_NUMBER,
			BuildNumber:   env.GI_APP_BUILD_NUMBER,
			BuildType:     env.GI_APP_BUILD_TYPE,
		},
	)

	fmt.Println()
	fmt.Println("Generating final icon....")
	fmt.Println("........................................")
	for index, iconFile := range baseIconsFile {
		fmt.Printf("Generating icon for base %d: %s\n", index, iconFile.Name())
		iconBase, _ := openIcon(env.GI_BASE_ICONS_SET + iconFile.Name())

		// Final icon combined both base + overlay
		icon := image.NewRGBA(iconBase.Bounds())
		draw.Draw(icon, icon.Bounds(), &image.Uniform{color.Black}, image.Point{0, 0}, draw.Over)
		draw.Draw(icon, icon.Bounds(), iconBase, image.Point{0, 0}, draw.Over)

		resizedOverlay := resize(iconOverlay, icon.Rect, env.GI_ICON_SCALER)
		draw.Draw(icon, icon.Bounds(), resizedOverlay, image.Point{0, 0}, draw.Over)

		// Saving final icon
		saveIcon(icon, env.GI_DESTINATION_ICONS_SET+iconFile.Name())
	}
	fmt.Println("Icons generated !!")
}

type (
	AppInfo struct {
		VersionNumber string
		BuildNumber   string
		BuildType     string
	}

	Overlay struct {
		Resolution     int
		Scale          int
		PrimaryColor   color.Color
		SecondaryColor color.Color
	}

	Label struct {
		Text  string
		Size  float64
		Color color.Color
		Rect  image.Rectangle
	}
)

func addLabel(icon *image.RGBA, label Label) {
	font, _ := freetype.ParseFont(goregular.TTF)

	c := freetype.NewContext()
	c.SetDPI(120)
	c.SetFont(font)
	c.SetFontSize(label.Size)
	c.SetClip(icon.Bounds())
	c.SetSrc(&image.Uniform{label.Color})

	// To measure total width, creating temp image
	measureBoard := image.NewRGBA(image.Rect(0, 0, 1, 1))
	c.SetDst(measureBoard)

	//positioning the label to center
	centerX := label.Rect.Min.X + label.Rect.Dx()/2
	centerY := label.Rect.Min.Y + label.Rect.Dy()/2 - int(label.Size) - 5
	pt := freetype.Pt(centerX, centerY+int(c.PointToFixed(label.Size)>>6))
	newPt, _ := c.DrawString(label.Text, pt)

	// Aligning text to center
	pt.X = pt.X - (newPt.X-pt.X)/2
	pt.Y = pt.Y - (newPt.Y-pt.Y)/2

	// Setting orignal icon
	c.SetDst(icon)
	c.DrawString(label.Text, pt)
}

func openIcon(path string) (image.Image, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	return img, err
}

func saveIcon(icon image.Image, path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	err = png.Encode(f, icon)
	if err != nil {
		return err
	}
	return nil
}

// https://stackoverflow.com/questions/22940724/go-resizing-images
func resize(icon image.Image, size image.Rectangle, scaler int) image.Image {
	resizedIcon := image.NewRGBA(size)
	if scaler == 1 {
		draw.ApproxBiLinear.Scale(resizedIcon, resizedIcon.Rect, icon, icon.Bounds(), draw.Over, nil)
	} else if scaler == 2 {
		draw.BiLinear.Scale(resizedIcon, resizedIcon.Rect, icon, icon.Bounds(), draw.Over, nil)
	} else if scaler == 3 {
		draw.CatmullRom.Scale(resizedIcon, resizedIcon.Rect, icon, icon.Bounds(), draw.Over, nil)
	} else {
		draw.NearestNeighbor.Scale(resizedIcon, resizedIcon.Rect, icon, icon.Bounds(), draw.Over, nil)
	}
	return resizedIcon
}

func createOverlay(oc Overlay, lc Label, ai AppInfo) image.Image {
	rect := image.Rect(0, 0, oc.Resolution, oc.Resolution)
	iconOverlay := image.NewRGBA(rect)

	zp := image.Point{}
	sectionHeight := rect.Dy() / 100 * oc.Scale
	sideGap := rect.Dx() / 100 * 20 / 2
	center := iconOverlay.Bounds().Dx() / 2

	versionRect := image.Rect(sideGap, 0, center, sectionHeight)
	buildRect := image.Rect(center, 0, iconOverlay.Rect.Dx()-sideGap, sectionHeight)
	titleRect := image.Rect(0, iconOverlay.Rect.Dy()-sectionHeight, iconOverlay.Rect.Dx(), iconOverlay.Rect.Dy())

	draw.Draw(iconOverlay, versionRect, &image.Uniform{oc.PrimaryColor}, zp, draw.Over)
	draw.Draw(iconOverlay, buildRect, &image.Uniform{oc.SecondaryColor}, zp, draw.Over)
	draw.Draw(iconOverlay, titleRect, &image.Uniform{oc.PrimaryColor}, zp, draw.Over)

	addLabel(iconOverlay, Label{Text: ai.VersionNumber, Color: lc.Color, Size: lc.Size, Rect: versionRect})
	addLabel(iconOverlay, Label{Text: ai.BuildNumber, Color: lc.Color, Size: lc.Size, Rect: buildRect})
	addLabel(iconOverlay, Label{Text: ai.BuildType, Color: lc.Color, Size: lc.Size, Rect: titleRect})
	return iconOverlay
}

func getFiles(dir string, regex *string) []os.DirEntry {
	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
		return []os.DirEntry{}
	}

	if regex != nil {
		matcher, _ := regexp.Compile(*regex)
		filteredFiles := []os.DirEntry{}
		for _, file := range files {
			if len(matcher.FindAllString(file.Name(), 1)) > 0 {
				filteredFiles = append(filteredFiles, file)
			}
		}
		return filteredFiles
	}
	return files
}
