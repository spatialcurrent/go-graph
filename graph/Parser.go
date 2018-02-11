package graph

type Parser interface{
  ParseQuery(q string) (OperationChain, error)
}
