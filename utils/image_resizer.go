package utils

//func CompressImage(fileHeader *multipart.FileHeader, compressionLevel int) (*multipart.FileHeader, error) {
//	file, err := fileHeader.Open()
//	if err != nil {
//		return nil, err
//	}
//	defer file.Close()
//
//	img, err := jpeg.Decode(file)
//	if err != nil {
//		return nil, err
//	}
//
//	// Perform image compression using resize package
//	// Adjust the quality value based on compression level
//	quality := uint(compressionLevel)
//	resizedImg := resize.Resize(0, 0, img, resize.MitchellNetravali)
//
//	// Create a buffer to store the compressed image
//	buf := new(bytes.Buffer)
//	err = jpeg.Encode(buf, resizedImg, &jpeg.Options{Quality: int(quality)})
//	if err != nil {
//		return nil, err
//	}
//
//	// Create a temporary file for the compressed image
//	tempFile, err := ioutil.TempFile("", "compressed_*.jpg")
//	if err != nil {
//		return nil, err
//	}
//	defer tempFile.Close()
//
//	// Write the compressed image data to the temporary file
//	_, err = tempFile.Write(buf.Bytes())
//	if err != nil {
//		return nil, err
//	}
//
//	// Create a new multipart.FileHeader for the compressed image
//	compressedFileHeader := &multipart.FileHeader{
//		Filename: fileHeader.Filename,
//		Size:     int64(buf.Len()),
//	}
//
//	return compressedFileHeader, nil
//}

//func CompressImageBy70Percent(fileHeader *multipart.FileHeader) (*multipart.FileHeader, error) {
//	file, err := fileHeader.Open()
//	if err != nil {
//		return nil, err
//	}
//	defer file.Close()
//
//	// Decode the image
//	img, err := jpeg.Decode(file)
//	if err != nil {
//		return nil, err
//	}
//
//	// Compress the image by 70%
//	newWidth := uint(float64(img.Bounds().Dx()) * 0.7)
//	newHeight := uint(float64(img.Bounds().Dy()) * 0.7)
//	resizedImg := resize.Resize(newWidth, newHeight, img, resize.Lanczos3)
//
//	// Create a temporary file to save the compressed image
//	tempFile, err := os.CreateTemp("", "compressed_image_*.jpg")
//	if err != nil {
//		return nil, err
//	}
//	defer tempFile.Close()
//
//	// Save the compressed image to the temporary file
//	err = jpeg.Encode(tempFile, resizedImg, nil)
//	if err != nil {
//		return nil, err
//	}
//
//	// Create a new *multipart.FileHeader for the compressed image
//	compressedFileHeader := &multipart.FileHeader{
//		Filename: fileHeader.Filename,
//		Size:     int64(tempFile.Size()),
//	}
//
//	// Open the temporary file for reading
//	compressedFile, err := tempFile.Open()
//	if err != nil {
//		return nil, err
//	}
//
//	// Set the *multipart.FileHeader's file field with the compressed file
//	compressedFileHeader.File = compressedFile
//
//	return compressedFileHeader, nil
//}
