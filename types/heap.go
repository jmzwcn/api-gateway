package types

type PrecisionSet []*MatchedMethod

func (pq *PrecisionSet) Max() *MatchedMethod {
	if pq == nil || *pq == nil {
		return nil
	}
	max := new(MatchedMethod)
	for i := 0; i < len(*pq); i++ {
		current := (*pq)[i]
		if max.Precision < current.Precision {
			max = current
		}
	}
	return max
}
