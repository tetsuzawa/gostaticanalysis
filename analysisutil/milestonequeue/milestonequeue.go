package milestonequeue

import "golang.org/x/tools/go/ssa"

type MilestoneQueue []ssa.Instruction

func (q *MilestoneQueue) push(f ssa.Instruction) () {
	*q = append(*q, f)
}

func (q *MilestoneQueue) pop() (f ssa.Instruction) {
	if q == nil {
		return
	}
	f = (*q)[0]
	*q = (*q)[1:]
	return f
}
