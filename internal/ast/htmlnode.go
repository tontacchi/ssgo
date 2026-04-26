package ast

type HTMLNode interface {
	ToHTML() (string, error)
}

