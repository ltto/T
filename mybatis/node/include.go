package node

type Include struct {
	RefId string
}

func (n *Include) Pare(args map[string]interface{}) (s string, err error) {
	sqlM := args["_sql"].(map[string]*DMLRoot)
	sql := sqlM[n.RefId]
	return PareNodes(args, sql.Child)
}
