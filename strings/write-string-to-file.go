package strings

import (
    "bufio"
    "os"
)

func WriteStringToFile(filename string, data string) error {
    // Open the file for writing, creating it if it doesn't exist
    file, err := os.Create(filename)
    if err != nil {
        return err
    }
    defer file.Close()

    // Create a new writer that writes to the file
    writer := bufio.NewWriter(file)

    // Write the string to the file
    _, err = writer.WriteString(data)
    if err != nil {
        return err
    }

    // Flush the writer to ensure all data is written to the file
    err = writer.Flush()
    if err != nil {
        return err
    }

    // Return nil to indicate success
    return nil
}
