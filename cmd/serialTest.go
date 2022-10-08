package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/tarm/serial"
)

// serialTestCmd represents the serialTest command
var serialTestCmd = &cobra.Command{
	Use:   "serial-test",
	Short: "Test the serial port",
	Long: `Test the serial port.
	
This is not a useful command it's just holding some test code for now.`,
	Run: func(cmd *cobra.Command, args []string) {
		c := &serial.Config{Name: "/dev/ttyUSB0", Baud: 57600}
		s, err := serial.OpenPort(c)
		if err != nil {
			log.Fatal(err)
		}

		n, err := s.Write([]byte("Q\x00\x00\xa0\x00\x00\x9d\x4ab"))
		if err != nil {
			log.Fatal(err)
		}

		buf := make([]byte, 128)
		n, err = s.Read(buf)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("%q", buf[:n])
	},
}

func init() {
	rootCmd.AddCommand(serialTestCmd)
}
