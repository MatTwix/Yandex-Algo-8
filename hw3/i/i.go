package main

import (
	"bufio"
	"os"
	"strings"
)

type TreeNode struct {
	Value string
	Left  *TreeNode
	Right *TreeNode
}

type Parser struct {
	input string
	pos   int
}

func NewParser(input string) *Parser {
	input = strings.ReplaceAll(input, " ", "")
	return &Parser{input: input, pos: 0}
}

func (p *Parser) peek() byte {
	if p.pos >= len(p.input) {
		return 0
	}
	return p.input[p.pos]
}

func (p *Parser) next() byte {
	if p.pos >= len(p.input) {
		return 0
	}
	ch := p.input[p.pos]
	p.pos++
	return ch
}

func (p *Parser) parseExpression() *TreeNode {
	left := p.parseTerm()

	for p.peek() == '+' || p.peek() == '-' {
		op := string(p.next())
		right := p.parseTerm()
		left = &TreeNode{Value: op, Left: left, Right: right}
	}

	return left
}

func (p *Parser) parseTerm() *TreeNode {
	left := p.parseFactor()

	for p.peek() == '*' || p.peek() == '/' {
		op := string(p.next())
		right := p.parseFactor()
		left = &TreeNode{Value: op, Left: left, Right: right}
	}

	return left
}

func (p *Parser) parseFactor() *TreeNode {
	left := p.parseElement()

	if p.peek() == '^' {
		p.next()
		right := p.parseFactor()
		return &TreeNode{Value: "^", Left: left, Right: right}
	}

	return left
}

func (p *Parser) parseElement() *TreeNode {
	ch := p.peek()

	if ch >= 'a' && ch <= 'z' {
		return &TreeNode{Value: string(p.next())}
	} else if ch == '(' {
		p.next()
		expr := p.parseExpression()
		p.next()
		return expr
	}

	return &TreeNode{}
}

func renderTree(node *TreeNode) []string {
	if node.Left == nil && node.Right == nil {
		return []string{node.Value}
	}

	leftTree := renderTree(node.Left)
	rightTree := renderTree(node.Right)

	lh := len(leftTree)
	rh := len(rightTree)
	lw := getWidth(leftTree)
	rw := getWidth(rightTree)
	height := max(lh, rh) + 2
	width := lw + rw + 5

	result := make([]string, height)
	for i := range result {
		result[i] = strings.Repeat(" ", width)
	}

	leftCenter := getTreeCenter(leftTree)
	rightCenter := lw + 5 + getTreeCenter(rightTree)

	rootCenter := lw + 3

	top := []rune(strings.Repeat(" ", width))
	if leftCenter >= 0 && leftCenter < len(top) {
		top[leftCenter] = '.'
	}
	if rightCenter >= 0 && rightCenter < len(top) {
		top[rightCenter] = '.'
	}
	for i := leftCenter + 1; i < rightCenter && i < len(top); i++ {
		if i >= 0 {
			top[i] = '-'
		}
	}

	if rootCenter-2 >= 0 && rootCenter < len(top) {
		top[rootCenter-2] = '['
		top[rootCenter-1] = rune(node.Value[0])
		top[rootCenter] = ']'
	} else {
		mid := (leftCenter + rightCenter) / 2
		if mid-2 >= 0 && mid < len(top) {
			top[mid-2] = '['
			top[mid-1] = rune(node.Value[0])
			top[mid] = ']'
		}
	}
	result[0] = string(top)

	sec := []rune(strings.Repeat(" ", width))
	if leftCenter >= 0 && leftCenter < len(sec) {
	
sec[leftCenter] = '|'
	}
	if rightCenter >= 0 && rightCenter < len(sec) {
	
sec[rightCenter] = '|'
	}
	result[1] = string(sec)

	for i, line := range leftTree {
		resRunes := []rune(result[2+i])
		copy(resRunes[:len(line)], []rune(line))
		result[2+i] = string(resRunes)
	}

	for i, line := range rightTree {
		resRunes := []rune(result[2+i])
		copy(resRunes[lw+5:], []rune(line))
		result[2+i] = string(resRunes)
	}

	return result
}

func getWidth(lines []string) int {
	maxWidth := 0
	for _, line := range lines {
		if len(line) > maxWidth {
			maxWidth = len(line)
		}
	}
	return maxWidth
}

func getTreeCenter(lines []string) int {
	if len(lines) == 1 {
		return 0
	}

	for i, ch := range lines[0] {
		if ch == '[' {
			return i + 1
		}
	}

	return 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)

	parser := NewParser(line)
	tree := parser.parseExpression()

	lines := renderTree(tree)

	for _, l := range lines {
		writer.WriteString(l)
		writer.WriteByte('\n')
	}
}
