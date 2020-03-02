package main

import "fmt"

func main() {
	testcases := []string{
		"[",
		"()",
		"()[]{}",
		"(]",
		"([)]",
		"{[]}",
	}
	for _, v := range testcases {
		fmt.Println(isValid(v))
	}
}

type StackNode struct {
	prev *StackNode
	data interface{}
}

type Stack struct {
	datas  *StackNode
	top    *StackNode
	bottom *StackNode
	height int
}

func (st *Stack) push(dt interface{}) {
	nNode := &StackNode{
		data: dt,
		prev: nil,
	}
	if st.bottom == nil || st.top == nil {
		st.bottom = nNode
		st.top = nNode
		st.datas = nNode
	} else {
		nNode.prev = st.top
		st.top = nNode
	}
	st.height += 1
}

func (st *Stack) pop() (data interface{}) {
	if st.top == nil || st.bottom == nil {
		return nil
	}
	if st.top != nil {
		data = st.top.data
		st.top = st.top.prev
		st.height -= 1
		return
	}
	return nil
}

func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}
	dtbyteTb := []byte{byte('('), byte(')'), byte('['), byte(']'), byte('{'), byte('}')}
	dtbyte := []byte(s)
	st := &Stack{}
	for _, v := range dtbyte {
		if v == dtbyteTb[0] || v == dtbyteTb[2] || v == dtbyteTb[4] {
			st.push(v)
		} else if v == dtbyteTb[1] || v == dtbyteTb[3] || v == dtbyteTb[5] {
			tmpDt := st.pop()
			if tmpDt != nil {
				switch v {
				case dtbyteTb[1]:
					if tmpDt != dtbyteTb[0] {
						return false
					}
				case dtbyteTb[3]:
					if tmpDt != dtbyteTb[2] {
						return false
					}
				case dtbyteTb[5]:
					if tmpDt != dtbyteTb[4] {
						return false
					}
				}
			} else {
				return false
			}
		} else {
			continue
		}
	}
	if st.height == 0 {
		return true
	}
	return false
}
