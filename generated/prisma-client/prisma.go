// Code generated by Prisma CLI (https://github.com/prisma/prisma). DO NOT EDIT.

package prisma

import (
	"context"
	"errors"

	"github.com/prisma/prisma-client-lib-go"

	"github.com/machinebox/graphql"
)

var ErrNoResult = errors.New("query returned no result")

func Str(v string) *string { return &v }
func Int32(v int32) *int32 { return &v }
func Bool(v bool) *bool    { return &v }

type BatchPayloadExec struct {
	exec *prisma.BatchPayloadExec
}

func (exec *BatchPayloadExec) Exec(ctx context.Context) (BatchPayload, error) {
	bp, err := exec.exec.Exec(ctx)
	return BatchPayload(bp), err
}

type BatchPayload struct {
	Count int64 `json:"count"`
}

type Aggregate struct {
	Count int64 `json:"count"`
}

type Client struct {
	Client *prisma.Client
}

type Options struct {
	Endpoint string
	Secret   string
}

func New(options *Options, opts ...graphql.ClientOption) *Client {
	endpoint := DefaultEndpoint
	secret := Secret
	if options != nil {
		endpoint = options.Endpoint
		secret = options.Secret
	}
	return &Client{
		Client: prisma.New(endpoint, secret, opts...),
	}
}

func (client *Client) GraphQL(ctx context.Context, query string, variables map[string]interface{}) (map[string]interface{}, error) {
	return client.Client.GraphQL(ctx, query, variables)
}

var DefaultEndpoint = "http://localhost:4466"
var Secret = ""

func (client *Client) Todo(params TodoWhereUniqueInput) *TodoExec {
	ret := client.Client.GetOne(
		nil,
		params,
		[2]string{"TodoWhereUniqueInput!", "Todo"},
		"todo",
		[]string{"id", "title", "completed", "body"})

	return &TodoExec{ret}
}

type TodoesParams struct {
	Where   *TodoWhereInput   `json:"where,omitempty"`
	OrderBy *TodoOrderByInput `json:"orderBy,omitempty"`
	Skip    *int32            `json:"skip,omitempty"`
	After   *string           `json:"after,omitempty"`
	Before  *string           `json:"before,omitempty"`
	First   *int32            `json:"first,omitempty"`
	Last    *int32            `json:"last,omitempty"`
}

func (client *Client) Todoes(params *TodoesParams) *TodoExecArray {
	var wparams *prisma.WhereParams
	if params != nil {
		wparams = &prisma.WhereParams{
			Where:   params.Where,
			OrderBy: (*string)(params.OrderBy),
			Skip:    params.Skip,
			After:   params.After,
			Before:  params.Before,
			First:   params.First,
			Last:    params.Last,
		}
	}

	ret := client.Client.GetMany(
		nil,
		wparams,
		[3]string{"TodoWhereInput", "TodoOrderByInput", "Todo"},
		"todoes",
		[]string{"id", "title", "completed", "body"})

	return &TodoExecArray{ret}
}

type TodoesConnectionParams struct {
	Where   *TodoWhereInput   `json:"where,omitempty"`
	OrderBy *TodoOrderByInput `json:"orderBy,omitempty"`
	Skip    *int32            `json:"skip,omitempty"`
	After   *string           `json:"after,omitempty"`
	Before  *string           `json:"before,omitempty"`
	First   *int32            `json:"first,omitempty"`
	Last    *int32            `json:"last,omitempty"`
}

// Nodes return just nodes without cursors. It uses the already fetched edges.
func (s *TodoConnection) Nodes() []Todo {
	var nodes []Todo
	for _, edge := range s.Edges {
		nodes = append(nodes, edge.Node)
	}
	return nodes
}

// Nodes return just nodes without cursors, but as a slice of pointers. It uses the already fetched edges.
func (s *TodoConnection) NodesPtr() []*Todo {
	var nodes []*Todo
	for _, edge := range s.Edges {
		item := edge
		nodes = append(nodes, &item.Node)
	}
	return nodes
}

func (client *Client) TodoesConnection(params *TodoesConnectionParams) *TodoConnectionExec {
	var wparams *prisma.WhereParams
	if params != nil {
		wparams = &prisma.WhereParams{
			Where:   params.Where,
			OrderBy: (*string)(params.OrderBy),
			Skip:    params.Skip,
			After:   params.After,
			Before:  params.Before,
			First:   params.First,
			Last:    params.Last,
		}
	}

	ret := client.Client.GetMany(
		nil,
		wparams,
		[3]string{"TodoWhereInput", "TodoOrderByInput", "Todo"},
		"todoesConnection",
		[]string{"edges", "pageInfo"})

	return &TodoConnectionExec{ret}
}

func (client *Client) CreateTodo(params TodoCreateInput) *TodoExec {
	ret := client.Client.Create(
		params,
		[2]string{"TodoCreateInput!", "Todo"},
		"createTodo",
		[]string{"id", "title", "completed", "body"})

	return &TodoExec{ret}
}

type TodoUpdateParams struct {
	Data  TodoUpdateInput      `json:"data"`
	Where TodoWhereUniqueInput `json:"where"`
}

func (client *Client) UpdateTodo(params TodoUpdateParams) *TodoExec {
	ret := client.Client.Update(
		prisma.UpdateParams{
			Data:  params.Data,
			Where: params.Where,
		},
		[3]string{"TodoUpdateInput!", "TodoWhereUniqueInput!", "Todo"},
		"updateTodo",
		[]string{"id", "title", "completed", "body"})

	return &TodoExec{ret}
}

type TodoUpdateManyParams struct {
	Data  TodoUpdateManyMutationInput `json:"data"`
	Where *TodoWhereInput             `json:"where,omitempty"`
}

func (client *Client) UpdateManyTodoes(params TodoUpdateManyParams) *BatchPayloadExec {
	exec := client.Client.UpdateMany(
		prisma.UpdateParams{
			Data:  params.Data,
			Where: params.Where,
		},
		[2]string{"TodoUpdateManyMutationInput!", "TodoWhereInput"},
		"updateManyTodoes")
	return &BatchPayloadExec{exec}
}

type TodoUpsertParams struct {
	Where  TodoWhereUniqueInput `json:"where"`
	Create TodoCreateInput      `json:"create"`
	Update TodoUpdateInput      `json:"update"`
}

func (client *Client) UpsertTodo(params TodoUpsertParams) *TodoExec {
	uparams := &prisma.UpsertParams{
		Where:  params.Where,
		Create: params.Create,
		Update: params.Update,
	}
	ret := client.Client.Upsert(
		uparams,
		[4]string{"TodoWhereUniqueInput!", "TodoCreateInput!", "TodoUpdateInput!", "Todo"},
		"upsertTodo",
		[]string{"id", "title", "completed", "body"})

	return &TodoExec{ret}
}

func (client *Client) DeleteTodo(params TodoWhereUniqueInput) *TodoExec {
	ret := client.Client.Delete(
		params,
		[2]string{"TodoWhereUniqueInput!", "Todo"},
		"deleteTodo",
		[]string{"id", "title", "completed", "body"})

	return &TodoExec{ret}
}

func (client *Client) DeleteManyTodoes(params *TodoWhereInput) *BatchPayloadExec {
	exec := client.Client.DeleteMany(params, "TodoWhereInput", "deleteManyTodoes")
	return &BatchPayloadExec{exec}
}

type TodoOrderByInput string

const (
	TodoOrderByInputIDAsc         TodoOrderByInput = "id_ASC"
	TodoOrderByInputIDDesc        TodoOrderByInput = "id_DESC"
	TodoOrderByInputTitleAsc      TodoOrderByInput = "title_ASC"
	TodoOrderByInputTitleDesc     TodoOrderByInput = "title_DESC"
	TodoOrderByInputCompletedAsc  TodoOrderByInput = "completed_ASC"
	TodoOrderByInputCompletedDesc TodoOrderByInput = "completed_DESC"
	TodoOrderByInputBodyAsc       TodoOrderByInput = "body_ASC"
	TodoOrderByInputBodyDesc      TodoOrderByInput = "body_DESC"
)

type MutationType string

const (
	MutationTypeCreated MutationType = "CREATED"
	MutationTypeUpdated MutationType = "UPDATED"
	MutationTypeDeleted MutationType = "DELETED"
)

type TodoWhereUniqueInput struct {
	ID *string `json:"id,omitempty"`
}

type TodoWhereInput struct {
	ID                 *string          `json:"id,omitempty"`
	IDNot              *string          `json:"id_not,omitempty"`
	IDIn               []string         `json:"id_in,omitempty"`
	IDNotIn            []string         `json:"id_not_in,omitempty"`
	IDLt               *string          `json:"id_lt,omitempty"`
	IDLte              *string          `json:"id_lte,omitempty"`
	IDGt               *string          `json:"id_gt,omitempty"`
	IDGte              *string          `json:"id_gte,omitempty"`
	IDContains         *string          `json:"id_contains,omitempty"`
	IDNotContains      *string          `json:"id_not_contains,omitempty"`
	IDStartsWith       *string          `json:"id_starts_with,omitempty"`
	IDNotStartsWith    *string          `json:"id_not_starts_with,omitempty"`
	IDEndsWith         *string          `json:"id_ends_with,omitempty"`
	IDNotEndsWith      *string          `json:"id_not_ends_with,omitempty"`
	Title              *string          `json:"title,omitempty"`
	TitleNot           *string          `json:"title_not,omitempty"`
	TitleIn            []string         `json:"title_in,omitempty"`
	TitleNotIn         []string         `json:"title_not_in,omitempty"`
	TitleLt            *string          `json:"title_lt,omitempty"`
	TitleLte           *string          `json:"title_lte,omitempty"`
	TitleGt            *string          `json:"title_gt,omitempty"`
	TitleGte           *string          `json:"title_gte,omitempty"`
	TitleContains      *string          `json:"title_contains,omitempty"`
	TitleNotContains   *string          `json:"title_not_contains,omitempty"`
	TitleStartsWith    *string          `json:"title_starts_with,omitempty"`
	TitleNotStartsWith *string          `json:"title_not_starts_with,omitempty"`
	TitleEndsWith      *string          `json:"title_ends_with,omitempty"`
	TitleNotEndsWith   *string          `json:"title_not_ends_with,omitempty"`
	Completed          *bool            `json:"completed,omitempty"`
	CompletedNot       *bool            `json:"completed_not,omitempty"`
	Body               *string          `json:"body,omitempty"`
	BodyNot            *string          `json:"body_not,omitempty"`
	BodyIn             []string         `json:"body_in,omitempty"`
	BodyNotIn          []string         `json:"body_not_in,omitempty"`
	BodyLt             *string          `json:"body_lt,omitempty"`
	BodyLte            *string          `json:"body_lte,omitempty"`
	BodyGt             *string          `json:"body_gt,omitempty"`
	BodyGte            *string          `json:"body_gte,omitempty"`
	BodyContains       *string          `json:"body_contains,omitempty"`
	BodyNotContains    *string          `json:"body_not_contains,omitempty"`
	BodyStartsWith     *string          `json:"body_starts_with,omitempty"`
	BodyNotStartsWith  *string          `json:"body_not_starts_with,omitempty"`
	BodyEndsWith       *string          `json:"body_ends_with,omitempty"`
	BodyNotEndsWith    *string          `json:"body_not_ends_with,omitempty"`
	And                []TodoWhereInput `json:"AND,omitempty"`
	Or                 []TodoWhereInput `json:"OR,omitempty"`
	Not                []TodoWhereInput `json:"NOT,omitempty"`
}

type TodoCreateInput struct {
	ID        *string `json:"id,omitempty"`
	Title     string  `json:"title"`
	Completed *bool   `json:"completed,omitempty"`
	Body      string  `json:"body"`
}

type TodoUpdateInput struct {
	Title     *string `json:"title,omitempty"`
	Completed *bool   `json:"completed,omitempty"`
	Body      *string `json:"body,omitempty"`
}

type TodoUpdateManyMutationInput struct {
	Title     *string `json:"title,omitempty"`
	Completed *bool   `json:"completed,omitempty"`
	Body      *string `json:"body,omitempty"`
}

type TodoSubscriptionWhereInput struct {
	MutationIn                 []MutationType               `json:"mutation_in,omitempty"`
	UpdatedFieldsContains      *string                      `json:"updatedFields_contains,omitempty"`
	UpdatedFieldsContainsEvery []string                     `json:"updatedFields_contains_every,omitempty"`
	UpdatedFieldsContainsSome  []string                     `json:"updatedFields_contains_some,omitempty"`
	Node                       *TodoWhereInput              `json:"node,omitempty"`
	And                        []TodoSubscriptionWhereInput `json:"AND,omitempty"`
	Or                         []TodoSubscriptionWhereInput `json:"OR,omitempty"`
	Not                        []TodoSubscriptionWhereInput `json:"NOT,omitempty"`
}

type TodoExec struct {
	exec *prisma.Exec
}

func (instance TodoExec) Exec(ctx context.Context) (*Todo, error) {
	var v Todo
	ok, err := instance.exec.Exec(ctx, &v)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, ErrNoResult
	}
	return &v, nil
}

func (instance TodoExec) Exists(ctx context.Context) (bool, error) {
	return instance.exec.Exists(ctx)
}

type TodoExecArray struct {
	exec *prisma.Exec
}

func (instance TodoExecArray) Exec(ctx context.Context) ([]Todo, error) {
	var v []Todo
	err := instance.exec.ExecArray(ctx, &v)
	return v, err
}

var TodoFields = []string{"id", "title", "completed", "body"}

type Todo struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
	Body      string `json:"body"`
}

type TodoConnectionExec struct {
	exec *prisma.Exec
}

func (instance *TodoConnectionExec) PageInfo() *PageInfoExec {
	ret := instance.exec.Client.GetOne(
		instance.exec,
		nil,
		[2]string{"", "PageInfo"},
		"pageInfo",
		[]string{"hasNextPage", "hasPreviousPage", "startCursor", "endCursor"})

	return &PageInfoExec{ret}
}

func (instance *TodoConnectionExec) Edges() *TodoEdgeExecArray {
	edges := instance.exec.Client.GetMany(
		instance.exec,
		nil,
		[3]string{"TodoWhereInput", "TodoOrderByInput", "TodoEdge"},
		"edges",
		[]string{"cursor"})

	nodes := edges.Client.GetOne(
		edges,
		nil,
		[2]string{"", "Todo"},
		"node",
		TodoFields)

	return &TodoEdgeExecArray{nodes}
}

func (instance *TodoConnectionExec) Aggregate(ctx context.Context) (*Aggregate, error) {
	ret := instance.exec.Client.GetOne(
		instance.exec,
		nil,
		[2]string{"", "AggregateTodo"},
		"aggregate",
		[]string{"count"})

	var v Aggregate
	_, err := ret.Exec(ctx, &v)
	return &v, err
}

func (instance TodoConnectionExec) Exec(ctx context.Context) (*TodoConnection, error) {
	edges, err := instance.Edges().Exec(ctx)
	if err != nil {
		return nil, err
	}

	pageInfo, err := instance.PageInfo().Exec(ctx)
	if err != nil {
		return nil, err
	}

	return &TodoConnection{
		Edges:    edges,
		PageInfo: *pageInfo,
	}, nil
}

func (instance TodoConnectionExec) Exists(ctx context.Context) (bool, error) {
	return instance.exec.Exists(ctx)
}

type TodoConnectionExecArray struct {
	exec *prisma.Exec
}

func (instance TodoConnectionExecArray) Exec(ctx context.Context) ([]TodoConnection, error) {
	var v []TodoConnection
	err := instance.exec.ExecArray(ctx, &v)
	return v, err
}

var TodoConnectionFields = []string{}

type TodoConnection struct {
	PageInfo PageInfo   `json:"pageInfo"`
	Edges    []TodoEdge `json:"edges"`
}

type PageInfoExec struct {
	exec *prisma.Exec
}

func (instance PageInfoExec) Exec(ctx context.Context) (*PageInfo, error) {
	var v PageInfo
	ok, err := instance.exec.Exec(ctx, &v)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, ErrNoResult
	}
	return &v, nil
}

func (instance PageInfoExec) Exists(ctx context.Context) (bool, error) {
	return instance.exec.Exists(ctx)
}

type PageInfoExecArray struct {
	exec *prisma.Exec
}

func (instance PageInfoExecArray) Exec(ctx context.Context) ([]PageInfo, error) {
	var v []PageInfo
	err := instance.exec.ExecArray(ctx, &v)
	return v, err
}

var PageInfoFields = []string{"hasNextPage", "hasPreviousPage", "startCursor", "endCursor"}

type PageInfo struct {
	HasNextPage     bool    `json:"hasNextPage"`
	HasPreviousPage bool    `json:"hasPreviousPage"`
	StartCursor     *string `json:"startCursor,omitempty"`
	EndCursor       *string `json:"endCursor,omitempty"`
}

type TodoEdgeExec struct {
	exec *prisma.Exec
}

func (instance *TodoEdgeExec) Node() *TodoExec {
	ret := instance.exec.Client.GetOne(
		instance.exec,
		nil,
		[2]string{"", "Todo"},
		"node",
		[]string{"id", "title", "completed", "body"})

	return &TodoExec{ret}
}

func (instance TodoEdgeExec) Exec(ctx context.Context) (*TodoEdge, error) {
	var v TodoEdge
	ok, err := instance.exec.Exec(ctx, &v)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, ErrNoResult
	}
	return &v, nil
}

func (instance TodoEdgeExec) Exists(ctx context.Context) (bool, error) {
	return instance.exec.Exists(ctx)
}

type TodoEdgeExecArray struct {
	exec *prisma.Exec
}

func (instance TodoEdgeExecArray) Exec(ctx context.Context) ([]TodoEdge, error) {
	var v []TodoEdge
	err := instance.exec.ExecArray(ctx, &v)
	return v, err
}

var TodoEdgeFields = []string{"cursor"}

type TodoEdge struct {
	Node   Todo   `json:"node"`
	Cursor string `json:"cursor"`
}

type TodoSubscriptionPayloadExec struct {
	exec *prisma.Exec
}

func (instance *TodoSubscriptionPayloadExec) Node() *TodoExec {
	ret := instance.exec.Client.GetOne(
		instance.exec,
		nil,
		[2]string{"", "Todo"},
		"node",
		[]string{"id", "title", "completed", "body"})

	return &TodoExec{ret}
}

func (instance *TodoSubscriptionPayloadExec) PreviousValues() *TodoPreviousValuesExec {
	ret := instance.exec.Client.GetOne(
		instance.exec,
		nil,
		[2]string{"", "TodoPreviousValues"},
		"previousValues",
		[]string{"id", "title", "completed", "body"})

	return &TodoPreviousValuesExec{ret}
}

func (instance TodoSubscriptionPayloadExec) Exec(ctx context.Context) (*TodoSubscriptionPayload, error) {
	var v TodoSubscriptionPayload
	ok, err := instance.exec.Exec(ctx, &v)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, ErrNoResult
	}
	return &v, nil
}

func (instance TodoSubscriptionPayloadExec) Exists(ctx context.Context) (bool, error) {
	return instance.exec.Exists(ctx)
}

type TodoSubscriptionPayloadExecArray struct {
	exec *prisma.Exec
}

func (instance TodoSubscriptionPayloadExecArray) Exec(ctx context.Context) ([]TodoSubscriptionPayload, error) {
	var v []TodoSubscriptionPayload
	err := instance.exec.ExecArray(ctx, &v)
	return v, err
}

var TodoSubscriptionPayloadFields = []string{"mutation", "updatedFields"}

type TodoSubscriptionPayload struct {
	Mutation      MutationType `json:"mutation"`
	Node          *Todo        `json:"node,omitempty"`
	UpdatedFields []string     `json:"updatedFields,omitempty"`
}

type TodoPreviousValuesExec struct {
	exec *prisma.Exec
}

func (instance TodoPreviousValuesExec) Exec(ctx context.Context) (*TodoPreviousValues, error) {
	var v TodoPreviousValues
	ok, err := instance.exec.Exec(ctx, &v)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, ErrNoResult
	}
	return &v, nil
}

func (instance TodoPreviousValuesExec) Exists(ctx context.Context) (bool, error) {
	return instance.exec.Exists(ctx)
}

type TodoPreviousValuesExecArray struct {
	exec *prisma.Exec
}

func (instance TodoPreviousValuesExecArray) Exec(ctx context.Context) ([]TodoPreviousValues, error) {
	var v []TodoPreviousValues
	err := instance.exec.ExecArray(ctx, &v)
	return v, err
}

var TodoPreviousValuesFields = []string{"id", "title", "completed", "body"}

type TodoPreviousValues struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
	Body      string `json:"body"`
}
