package milestonequeue

import "go/types"

type MilestoneQueue []types.Object

func (q *MilestoneQueue) Push(t types.Object) {
	*q = append(*q, t)
}

func (q *MilestoneQueue) Pop() (t types.Object) {
	if q == nil {
		return nil
	}
	t = (*q)[0]
	*q = (*q)[1:]
	return t
}

func (q *MilestoneQueue) Head() types.Object {
	if q.Len() == 0 {
		return nil
	}
	return (*q)[0]
}

func (q *MilestoneQueue) Tail() types.Object {
	if q.Len() == 0 {
		return nil
	}
	return (*q)[(*q).Len()-1]
}

func (q *MilestoneQueue) Len() int {
	return len(*q)
}
