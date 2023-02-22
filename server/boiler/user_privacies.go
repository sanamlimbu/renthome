// Code generated by SQLBoiler 4.14.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package boiler

import (
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// UserPrivacy is an object representing the database table.
type UserPrivacy struct {
	UserID    string    `boiler:"user_id" boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	PrivacyID string    `boiler:"privacy_id" boil:"privacy_id" json:"privacy_id" toml:"privacy_id" yaml:"privacy_id"`
	State     string    `boiler:"state" boil:"state" json:"state" toml:"state" yaml:"state"`
	CreatedAt time.Time `boiler:"created_at" boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt time.Time `boiler:"updated_at" boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *userPrivacyR `boiler:"-" boil:"-" json:"-" toml:"-" yaml:"-"`
	L userPrivacyL  `boiler:"-" boil:"-" json:"-" toml:"-" yaml:"-"`
}

var UserPrivacyColumns = struct {
	UserID    string
	PrivacyID string
	State     string
	CreatedAt string
	UpdatedAt string
}{
	UserID:    "user_id",
	PrivacyID: "privacy_id",
	State:     "state",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

var UserPrivacyTableColumns = struct {
	UserID    string
	PrivacyID string
	State     string
	CreatedAt string
	UpdatedAt string
}{
	UserID:    "user_privacies.user_id",
	PrivacyID: "user_privacies.privacy_id",
	State:     "user_privacies.state",
	CreatedAt: "user_privacies.created_at",
	UpdatedAt: "user_privacies.updated_at",
}

// Generated where

var UserPrivacyWhere = struct {
	UserID    whereHelperstring
	PrivacyID whereHelperstring
	State     whereHelperstring
	CreatedAt whereHelpertime_Time
	UpdatedAt whereHelpertime_Time
}{
	UserID:    whereHelperstring{field: "\"user_privacies\".\"user_id\""},
	PrivacyID: whereHelperstring{field: "\"user_privacies\".\"privacy_id\""},
	State:     whereHelperstring{field: "\"user_privacies\".\"state\""},
	CreatedAt: whereHelpertime_Time{field: "\"user_privacies\".\"created_at\""},
	UpdatedAt: whereHelpertime_Time{field: "\"user_privacies\".\"updated_at\""},
}

// UserPrivacyRels is where relationship names are stored.
var UserPrivacyRels = struct {
	Privacy string
	User    string
}{
	Privacy: "Privacy",
	User:    "User",
}

// userPrivacyR is where relationships are stored.
type userPrivacyR struct {
	Privacy *Privacy `boiler:"Privacy" boil:"Privacy" json:"Privacy" toml:"Privacy" yaml:"Privacy"`
	User    *User    `boiler:"User" boil:"User" json:"User" toml:"User" yaml:"User"`
}

// NewStruct creates a new relationship struct
func (*userPrivacyR) NewStruct() *userPrivacyR {
	return &userPrivacyR{}
}

func (r *userPrivacyR) GetPrivacy() *Privacy {
	if r == nil {
		return nil
	}
	return r.Privacy
}

func (r *userPrivacyR) GetUser() *User {
	if r == nil {
		return nil
	}
	return r.User
}

// userPrivacyL is where Load methods for each relationship are stored.
type userPrivacyL struct{}

var (
	userPrivacyAllColumns            = []string{"user_id", "privacy_id", "state", "created_at", "updated_at"}
	userPrivacyColumnsWithoutDefault = []string{"user_id", "privacy_id"}
	userPrivacyColumnsWithDefault    = []string{"state", "created_at", "updated_at"}
	userPrivacyPrimaryKeyColumns     = []string{"user_id", "privacy_id"}
	userPrivacyGeneratedColumns      = []string{}
)

type (
	// UserPrivacySlice is an alias for a slice of pointers to UserPrivacy.
	// This should almost always be used instead of []UserPrivacy.
	UserPrivacySlice []*UserPrivacy
	// UserPrivacyHook is the signature for custom UserPrivacy hook methods
	UserPrivacyHook func(boil.Executor, *UserPrivacy) error

	userPrivacyQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	userPrivacyType                 = reflect.TypeOf(&UserPrivacy{})
	userPrivacyMapping              = queries.MakeStructMapping(userPrivacyType)
	userPrivacyPrimaryKeyMapping, _ = queries.BindMapping(userPrivacyType, userPrivacyMapping, userPrivacyPrimaryKeyColumns)
	userPrivacyInsertCacheMut       sync.RWMutex
	userPrivacyInsertCache          = make(map[string]insertCache)
	userPrivacyUpdateCacheMut       sync.RWMutex
	userPrivacyUpdateCache          = make(map[string]updateCache)
	userPrivacyUpsertCacheMut       sync.RWMutex
	userPrivacyUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var userPrivacyAfterSelectHooks []UserPrivacyHook

var userPrivacyBeforeInsertHooks []UserPrivacyHook
var userPrivacyAfterInsertHooks []UserPrivacyHook

var userPrivacyBeforeUpdateHooks []UserPrivacyHook
var userPrivacyAfterUpdateHooks []UserPrivacyHook

var userPrivacyBeforeDeleteHooks []UserPrivacyHook
var userPrivacyAfterDeleteHooks []UserPrivacyHook

var userPrivacyBeforeUpsertHooks []UserPrivacyHook
var userPrivacyAfterUpsertHooks []UserPrivacyHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *UserPrivacy) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range userPrivacyAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *UserPrivacy) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range userPrivacyBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *UserPrivacy) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range userPrivacyAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *UserPrivacy) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range userPrivacyBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *UserPrivacy) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range userPrivacyAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *UserPrivacy) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range userPrivacyBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *UserPrivacy) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range userPrivacyAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *UserPrivacy) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range userPrivacyBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *UserPrivacy) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range userPrivacyAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddUserPrivacyHook registers your hook function for all future operations.
func AddUserPrivacyHook(hookPoint boil.HookPoint, userPrivacyHook UserPrivacyHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		userPrivacyAfterSelectHooks = append(userPrivacyAfterSelectHooks, userPrivacyHook)
	case boil.BeforeInsertHook:
		userPrivacyBeforeInsertHooks = append(userPrivacyBeforeInsertHooks, userPrivacyHook)
	case boil.AfterInsertHook:
		userPrivacyAfterInsertHooks = append(userPrivacyAfterInsertHooks, userPrivacyHook)
	case boil.BeforeUpdateHook:
		userPrivacyBeforeUpdateHooks = append(userPrivacyBeforeUpdateHooks, userPrivacyHook)
	case boil.AfterUpdateHook:
		userPrivacyAfterUpdateHooks = append(userPrivacyAfterUpdateHooks, userPrivacyHook)
	case boil.BeforeDeleteHook:
		userPrivacyBeforeDeleteHooks = append(userPrivacyBeforeDeleteHooks, userPrivacyHook)
	case boil.AfterDeleteHook:
		userPrivacyAfterDeleteHooks = append(userPrivacyAfterDeleteHooks, userPrivacyHook)
	case boil.BeforeUpsertHook:
		userPrivacyBeforeUpsertHooks = append(userPrivacyBeforeUpsertHooks, userPrivacyHook)
	case boil.AfterUpsertHook:
		userPrivacyAfterUpsertHooks = append(userPrivacyAfterUpsertHooks, userPrivacyHook)
	}
}

// One returns a single userPrivacy record from the query.
func (q userPrivacyQuery) One(exec boil.Executor) (*UserPrivacy, error) {
	o := &UserPrivacy{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(nil, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "boiler: failed to execute a one query for user_privacies")
	}

	if err := o.doAfterSelectHooks(exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all UserPrivacy records from the query.
func (q userPrivacyQuery) All(exec boil.Executor) (UserPrivacySlice, error) {
	var o []*UserPrivacy

	err := q.Bind(nil, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "boiler: failed to assign all query results to UserPrivacy slice")
	}

	if len(userPrivacyAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all UserPrivacy records in the query.
func (q userPrivacyQuery) Count(exec boil.Executor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: failed to count user_privacies rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q userPrivacyQuery) Exists(exec boil.Executor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "boiler: failed to check if user_privacies exists")
	}

	return count > 0, nil
}

// Privacy pointed to by the foreign key.
func (o *UserPrivacy) Privacy(mods ...qm.QueryMod) privacyQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.PrivacyID),
	}

	queryMods = append(queryMods, mods...)

	return Privacies(queryMods...)
}

// User pointed to by the foreign key.
func (o *UserPrivacy) User(mods ...qm.QueryMod) userQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.UserID),
	}

	queryMods = append(queryMods, mods...)

	return Users(queryMods...)
}

// LoadPrivacy allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (userPrivacyL) LoadPrivacy(e boil.Executor, singular bool, maybeUserPrivacy interface{}, mods queries.Applicator) error {
	var slice []*UserPrivacy
	var object *UserPrivacy

	if singular {
		var ok bool
		object, ok = maybeUserPrivacy.(*UserPrivacy)
		if !ok {
			object = new(UserPrivacy)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeUserPrivacy)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeUserPrivacy))
			}
		}
	} else {
		s, ok := maybeUserPrivacy.(*[]*UserPrivacy)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeUserPrivacy)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeUserPrivacy))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &userPrivacyR{}
		}
		args = append(args, object.PrivacyID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &userPrivacyR{}
			}

			for _, a := range args {
				if a == obj.PrivacyID {
					continue Outer
				}
			}

			args = append(args, obj.PrivacyID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`privacies`),
		qm.WhereIn(`privacies.id in ?`, args...),
		qmhelper.WhereIsNull(`privacies.deleted_at`),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.Query(e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Privacy")
	}

	var resultSlice []*Privacy
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Privacy")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for privacies")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for privacies")
	}

	if len(privacyAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Privacy = foreign
		if foreign.R == nil {
			foreign.R = &privacyR{}
		}
		foreign.R.UserPrivacies = append(foreign.R.UserPrivacies, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.PrivacyID == foreign.ID {
				local.R.Privacy = foreign
				if foreign.R == nil {
					foreign.R = &privacyR{}
				}
				foreign.R.UserPrivacies = append(foreign.R.UserPrivacies, local)
				break
			}
		}
	}

	return nil
}

// LoadUser allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (userPrivacyL) LoadUser(e boil.Executor, singular bool, maybeUserPrivacy interface{}, mods queries.Applicator) error {
	var slice []*UserPrivacy
	var object *UserPrivacy

	if singular {
		var ok bool
		object, ok = maybeUserPrivacy.(*UserPrivacy)
		if !ok {
			object = new(UserPrivacy)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeUserPrivacy)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeUserPrivacy))
			}
		}
	} else {
		s, ok := maybeUserPrivacy.(*[]*UserPrivacy)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeUserPrivacy)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeUserPrivacy))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &userPrivacyR{}
		}
		args = append(args, object.UserID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &userPrivacyR{}
			}

			for _, a := range args {
				if a == obj.UserID {
					continue Outer
				}
			}

			args = append(args, obj.UserID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`users`),
		qm.WhereIn(`users.id in ?`, args...),
		qmhelper.WhereIsNull(`users.deleted_at`),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.Query(e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load User")
	}

	var resultSlice []*User
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice User")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for users")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for users")
	}

	if len(userAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.User = foreign
		if foreign.R == nil {
			foreign.R = &userR{}
		}
		foreign.R.UserPrivacies = append(foreign.R.UserPrivacies, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.UserID == foreign.ID {
				local.R.User = foreign
				if foreign.R == nil {
					foreign.R = &userR{}
				}
				foreign.R.UserPrivacies = append(foreign.R.UserPrivacies, local)
				break
			}
		}
	}

	return nil
}

// SetPrivacy of the userPrivacy to the related item.
// Sets o.R.Privacy to related.
// Adds o to related.R.UserPrivacies.
func (o *UserPrivacy) SetPrivacy(exec boil.Executor, insert bool, related *Privacy) error {
	var err error
	if insert {
		if err = related.Insert(exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"user_privacies\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"privacy_id"}),
		strmangle.WhereClause("\"", "\"", 2, userPrivacyPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.UserID, o.PrivacyID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}
	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.PrivacyID = related.ID
	if o.R == nil {
		o.R = &userPrivacyR{
			Privacy: related,
		}
	} else {
		o.R.Privacy = related
	}

	if related.R == nil {
		related.R = &privacyR{
			UserPrivacies: UserPrivacySlice{o},
		}
	} else {
		related.R.UserPrivacies = append(related.R.UserPrivacies, o)
	}

	return nil
}

// SetUser of the userPrivacy to the related item.
// Sets o.R.User to related.
// Adds o to related.R.UserPrivacies.
func (o *UserPrivacy) SetUser(exec boil.Executor, insert bool, related *User) error {
	var err error
	if insert {
		if err = related.Insert(exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"user_privacies\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"user_id"}),
		strmangle.WhereClause("\"", "\"", 2, userPrivacyPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.UserID, o.PrivacyID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}
	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.UserID = related.ID
	if o.R == nil {
		o.R = &userPrivacyR{
			User: related,
		}
	} else {
		o.R.User = related
	}

	if related.R == nil {
		related.R = &userR{
			UserPrivacies: UserPrivacySlice{o},
		}
	} else {
		related.R.UserPrivacies = append(related.R.UserPrivacies, o)
	}

	return nil
}

// UserPrivacies retrieves all the records using an executor.
func UserPrivacies(mods ...qm.QueryMod) userPrivacyQuery {
	mods = append(mods, qm.From("\"user_privacies\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"user_privacies\".*"})
	}

	return userPrivacyQuery{q}
}

// FindUserPrivacy retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindUserPrivacy(exec boil.Executor, userID string, privacyID string, selectCols ...string) (*UserPrivacy, error) {
	userPrivacyObj := &UserPrivacy{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"user_privacies\" where \"user_id\"=$1 AND \"privacy_id\"=$2", sel,
	)

	q := queries.Raw(query, userID, privacyID)

	err := q.Bind(nil, exec, userPrivacyObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "boiler: unable to select from user_privacies")
	}

	if err = userPrivacyObj.doAfterSelectHooks(exec); err != nil {
		return userPrivacyObj, err
	}

	return userPrivacyObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *UserPrivacy) Insert(exec boil.Executor, columns boil.Columns) error {
	if o == nil {
		return errors.New("boiler: no user_privacies provided for insertion")
	}

	var err error
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}
	if o.UpdatedAt.IsZero() {
		o.UpdatedAt = currTime
	}

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(userPrivacyColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	userPrivacyInsertCacheMut.RLock()
	cache, cached := userPrivacyInsertCache[key]
	userPrivacyInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			userPrivacyAllColumns,
			userPrivacyColumnsWithDefault,
			userPrivacyColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(userPrivacyType, userPrivacyMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(userPrivacyType, userPrivacyMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"user_privacies\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"user_privacies\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRow(cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.Exec(cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "boiler: unable to insert into user_privacies")
	}

	if !cached {
		userPrivacyInsertCacheMut.Lock()
		userPrivacyInsertCache[key] = cache
		userPrivacyInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// Update uses an executor to update the UserPrivacy.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *UserPrivacy) Update(exec boil.Executor, columns boil.Columns) (int64, error) {
	currTime := time.Now().In(boil.GetLocation())

	o.UpdatedAt = currTime

	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	userPrivacyUpdateCacheMut.RLock()
	cache, cached := userPrivacyUpdateCache[key]
	userPrivacyUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			userPrivacyAllColumns,
			userPrivacyPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("boiler: unable to update user_privacies, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"user_privacies\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, userPrivacyPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(userPrivacyType, userPrivacyMapping, append(wl, userPrivacyPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}
	var result sql.Result
	result, err = exec.Exec(cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to update user_privacies row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: failed to get rows affected by update for user_privacies")
	}

	if !cached {
		userPrivacyUpdateCacheMut.Lock()
		userPrivacyUpdateCache[key] = cache
		userPrivacyUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(exec)
}

// UpdateAll updates all rows with the specified column values.
func (q userPrivacyQuery) UpdateAll(exec boil.Executor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to update all for user_privacies")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to retrieve rows affected for user_privacies")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o UserPrivacySlice) UpdateAll(exec boil.Executor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("boiler: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), userPrivacyPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"user_privacies\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, userPrivacyPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to update all in userPrivacy slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to retrieve rows affected all in update all userPrivacy")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *UserPrivacy) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("boiler: no user_privacies provided for upsert")
	}
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}
	o.UpdatedAt = currTime

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(userPrivacyColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	userPrivacyUpsertCacheMut.RLock()
	cache, cached := userPrivacyUpsertCache[key]
	userPrivacyUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			userPrivacyAllColumns,
			userPrivacyColumnsWithDefault,
			userPrivacyColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			userPrivacyAllColumns,
			userPrivacyPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("boiler: unable to upsert user_privacies, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(userPrivacyPrimaryKeyColumns))
			copy(conflict, userPrivacyPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"user_privacies\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(userPrivacyType, userPrivacyMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(userPrivacyType, userPrivacyMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRow(cache.query, vals...).Scan(returns...)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.Exec(cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "boiler: unable to upsert user_privacies")
	}

	if !cached {
		userPrivacyUpsertCacheMut.Lock()
		userPrivacyUpsertCache[key] = cache
		userPrivacyUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// Delete deletes a single UserPrivacy record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *UserPrivacy) Delete(exec boil.Executor) (int64, error) {
	if o == nil {
		return 0, errors.New("boiler: no UserPrivacy provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), userPrivacyPrimaryKeyMapping)
	sql := "DELETE FROM \"user_privacies\" WHERE \"user_id\"=$1 AND \"privacy_id\"=$2"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to delete from user_privacies")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: failed to get rows affected by delete for user_privacies")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q userPrivacyQuery) DeleteAll(exec boil.Executor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("boiler: no userPrivacyQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to delete all from user_privacies")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: failed to get rows affected by deleteall for user_privacies")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o UserPrivacySlice) DeleteAll(exec boil.Executor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(userPrivacyBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), userPrivacyPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"user_privacies\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, userPrivacyPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to delete all from userPrivacy slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: failed to get rows affected by deleteall for user_privacies")
	}

	if len(userPrivacyAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *UserPrivacy) Reload(exec boil.Executor) error {
	ret, err := FindUserPrivacy(exec, o.UserID, o.PrivacyID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *UserPrivacySlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := UserPrivacySlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), userPrivacyPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"user_privacies\".* FROM \"user_privacies\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, userPrivacyPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(nil, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "boiler: unable to reload all in UserPrivacySlice")
	}

	*o = slice

	return nil
}

// UserPrivacyExists checks if the UserPrivacy row exists.
func UserPrivacyExists(exec boil.Executor, userID string, privacyID string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"user_privacies\" where \"user_id\"=$1 AND \"privacy_id\"=$2 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, userID, privacyID)
	}
	row := exec.QueryRow(sql, userID, privacyID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "boiler: unable to check if user_privacies exists")
	}

	return exists, nil
}

// Exists checks if the UserPrivacy row exists.
func (o *UserPrivacy) Exists(exec boil.Executor) (bool, error) {
	return UserPrivacyExists(exec, o.UserID, o.PrivacyID)
}