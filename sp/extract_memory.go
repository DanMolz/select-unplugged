package sp

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
		return Memory{
			area: area,
			data: memory.Data()[start:end],
		}
	}
	panic("Area not found")
}
