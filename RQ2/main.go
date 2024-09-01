package main

import vdfs "github.com/shreekarashastry/vdf-simulation/RQ2/vdfs"

const vdf_type = 1

func main() {
	if vdf_type == 0 {
		vdfs.RunWesolowski()
	} else if vdf_type == 1 {
		vdfs.RunPietrzaks()
	}
}
