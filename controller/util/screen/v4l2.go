package screen

import (
	"context"
	"fmt"
	"github.com/vladimirvivien/go4vl/device"
	"github.com/vladimirvivien/go4vl/v4l2"
	"log"
	"os"
)

func Video() {
	devName := "/dev/video0"
	device, err := device.Open(
		devName,
		device.WithPixFormat(v4l2.PixFormat{PixelFormat: v4l2.PixelFmtMPEG, Width: 640, Height: 480}),
	)
	if err != nil {
		log.Fatalf("failed to open device: %s", err)
	}
	defer device.Close()
	// start stream with cancellable context
	ctx, stop := context.WithCancel(context.TODO())
	if err := device.Start(ctx); err != nil {
		log.Fatalf("failed to start stream: %s", err)
	}
	// process frames from capture channel
	totalFrames := 10
	log.Printf("Capturing %d frames...", totalFrames)
	fileName := "V-2af608bf-80c3-451d-b5b9-05b6c36d53c9.mp4"

	for frame := range device.GetOutput() {
		file, err := os.Create(devName)
		if err != nil {
			log.Printf("failed to create file %s: %s", devName, err)
		}
		if _, err := file.Write(frame); err != nil {
			log.Printf("failed to write file %s: %s", fileName, err)
			continue
		}
		log.Printf("Saved file: %s", fileName)
		if err := file.Close(); err != nil {
			log.Printf("failed to close file %s: %s", fileName, err)
		}
		break
	}
	stop()
	fmt.Println("Done.")
}
