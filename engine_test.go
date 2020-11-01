package main

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type EngineTestSuite struct {
	suite.Suite
	nw *HirschbergSolver
}

func (s *EngineTestSuite) TestDNA() {
	for _, test := range []struct {
		seq1    string
		seq2    string
		result1 string
		result2 string
		score   int
	}{
		{
			"AATCG",
			"AACG",
			"AATCG",
			"AA-CG",
			10,
		},
		{
			"GTACAACGTTA",
			"AATCGTAGCGA",
			"GTA-CAACGTTA",
			"A-ATCGTAGCGA",
			-24,
		},
		{
			"GCGCGTGCGCGGAAGGAGCCAAGGTGAAGTTGTAGCAGTGTGTCAGAAGAGGTGCGTGGCACCATGCTGTCCCCCGAGGCGGAGCGGGTGCTGCGGTACCTGGTCGAAGTAGAGGAGTTG",
			"GACTTGTGGAACCTACTTCCTGAAAATAACCTTCTGTCCTCCGAGCTCTCCGCACCCGTGGATGACCTGCTCCCGTACACAGATGTTGCCACCTGGCTGGATGAATGTCCGAATGAAGCG",
			"GCG-CGTGCGCGGAAGGAGCCAAGGTGAAGTTGTAGCAGTGTGTCAGAAGAGGTGCGT-GGCACCA-TGC-TGTCC--C-CC-GAGG--CGGAGCGGGTGCTGCGGTACCTGGTCGAA-GTA-GA--GG-AGTTG",
			"G--AC-T-TGTGGAA-CCTACTTCCTGAAAAT--AACCTTCTGTCCTCCGAGCT-CT-CCGCACCCGTGGATGACCTGCTCCCG---TACACAGATGTTGCCAC-CTGGCTGGATGAATGTCCGAATG-AAG-CG",
			-135,
		},
	} {
		s.nw = Hirschberg(&Sequence{Value: test.seq1}, &Sequence{Value: test.seq2}, DNAFull, -10)
		r1, r2, sc := s.nw.Solve()
		s.Equal(r1, test.result1)
		s.Equal(r2, test.result2)
		s.Equal(sc, test.score)
	}
}

func TestLab2Suite(t *testing.T) {
	suite.Run(t, new(EngineTestSuite))
}
