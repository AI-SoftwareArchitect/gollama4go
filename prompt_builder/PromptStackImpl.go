package prompt_builder

import "fmt"

func (s *PromptStack) Push(prompt string) {
	s.PromptData = append(s.PromptData, prompt)
}

func (s *PromptStack) Pop() (string, bool) {
	if len(s.PromptData) > 0 {
		last := s.PromptData[len(s.PromptData)-1]
		s.PromptData = s.PromptData[:len(s.PromptData)-1]
		return last, true
	}
	return "", false
}

func (s *PromptStack) Peek() (string, bool) {
	if len(s.PromptData) > 0 {
		lastIndx := len(s.PromptData) - 1
		return s.PromptData[lastIndx], true
	}
	return "", false
}

func (s *PromptStack) Print() {
	for _, line := range s.PromptData {
		fmt.Println(line)
	}
}

func (s *PromptStack) Len() int {
	return len(s.PromptData)
}

func (s *PromptStack) Empty() bool {
	return len(s.PromptData) == 0
}
