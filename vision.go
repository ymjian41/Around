package main

import (
        "context"
        "fmt"

        vision "cloud.google.com/go/vision/apiv1"
)

// Annotate an image file based on Cloud Vision API, return score and error if exists.
func annotate(uri string) (float32, error) {

        ctx := context.Background()
        client, err := vision.NewImageAnnotatorClient(ctx)
        if err != nil {
                return 0.0, nil
        }
        defer client.Close()

        image := vision.NewImageFromURI(uri)
        annotations, err := client.DetectFaces(ctx, image, nil, 1)
        if err != nil {
                return 0.0, nil
        }

        if len(annotations) == 0 {
                fmt.Println("No faces found.")
                return 0.0, nil
        }

        return annotations[0].DetectionConfidence, nil
}
