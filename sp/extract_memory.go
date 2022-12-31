package sp

import "log"

func ExtractMemory(area Area, memories []Memory) Memory {

	for i := 0; i < len(memories); i++ {
		memory := memories[i]
		start := (area.address - memory.Area().address) * 2
		if start < 0 {
			continue
		}
		end := start + Address(area.Words()*2)
		if end > Address(memory.Area().Words()*2) {
			continue
		}
		log.Printf("start: %d, end: %d", start, end)
		return Memory{
			area: area,
			data: memory.Data()[start:end],
		}
	}
	panic("Area not found")
}
