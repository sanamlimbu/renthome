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
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// ResetPassword is an object representing the database table.
type ResetPassword struct {
	UserID    string    `boiler:"user_id" boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	Code      string    `boiler:"code" boil:"code" json:"code" toml:"code" yaml:"code"`
	UpdatedAt time.Time `boiler:"updated_at" boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`
	CreatedAt time.Time `boiler:"created_at" boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	DeletedAt null.Time `boiler:"deleted_at" boil:"deleted_at" json:"deleted_at,omitempty" toml:"deleted_at" yaml:"deleted_at,omitempty"`

	R *resetPasswordR `boiler:"-" boil:"-" json:"-" toml:"-" yaml:"-"`
	L resetPasswordL  `boiler:"-" boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ResetPasswordColumns = struct {
	UserID    string
	Code      string
	UpdatedAt string
	CreatedAt string
	DeletedAt string
}{
	UserID:    "user_id",
	Code:      "code",
	UpdatedAt: "updated_at",
	CreatedAt: "created_at",
	DeletedAt: "deleted_at",
}

var ResetPasswordTableColumns = struct {
	UserID    string
	Code      string
	UpdatedAt string
	CreatedAt string
	DeletedAt string
}{
	UserID:    "reset_passwords.user_id",
	Code:      "reset_passwords.code",
	UpdatedAt: "reset_passwords.updated_at",
	CreatedAt: "reset_passwords.created_at",
	DeletedAt: "reset_passwords.deleted_at",
}

// Generated where

var ResetPasswordWhere = struct {
	UserID    whereHelperstring
	Code      whereHelperstring
	UpdatedAt whereHelpertime_Time
	CreatedAt whereHelpertime_Time
	DeletedAt whereHelpernull_Time
}{
	UserID:    whereHelperstring{field: "\"reset_passwords\".\"user_id\""},
	Code:      whereHelperstring{field: "\"reset_passwords\".\"code\""},
	UpdatedAt: whereHelpertime_Time{field: "\"reset_passwords\".\"updated_at\""},
	CreatedAt: whereHelpertime_Time{field: "\"reset_passwords\".\"created_at\""},
	DeletedAt: whereHelpernull_Time{field: "\"reset_passwords\".\"deleted_at\""},
}

// ResetPasswordRels is where relationship names are stored.
var ResetPasswordRels = struct {
	User string
}{
	User: "User",
}

// resetPasswordR is where relationships are stored.
type resetPasswordR struct {
	User *User `boiler:"User" boil:"User" json:"User" toml:"User" yaml:"User"`
}

// NewStruct creates a new relationship struct
func (*resetPasswordR) NewStruct() *resetPasswordR {
	return &resetPasswordR{}
}

func (r *resetPasswordR) GetUser() *User {
	if r == nil {
		return nil
	}
	return r.User
}

// resetPasswordL is where Load methods for each relationship are stored.
type resetPasswordL struct{}

var (
	resetPasswordAllColumns            = []string{"user_id", "code", "updated_at", "created_at", "deleted_at"}
	resetPasswordColumnsWithoutDefault = []string{"user_id", "code"}
	resetPasswordColumnsWithDefault    = []string{"updated_at", "created_at", "deleted_at"}
	resetPasswordPrimaryKeyColumns     = []string{"user_id"}
	resetPasswordGeneratedColumns      = []string{}
)

type (
	// ResetPasswordSlice is an alias for a slice of pointers to ResetPassword.
	// This should almost always be used instead of []ResetPassword.
	ResetPasswordSlice []*ResetPassword
	// ResetPasswordHook is the signature for custom ResetPassword hook methods
	ResetPasswordHook func(boil.Executor, *ResetPassword) error

	resetPasswordQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	resetPasswordType                 = reflect.TypeOf(&ResetPassword{})
	resetPasswordMapping              = queries.MakeStructMapping(resetPasswordType)
	resetPasswordPrimaryKeyMapping, _ = queries.BindMapping(resetPasswordType, resetPasswordMapping, resetPasswordPrimaryKeyColumns)
	resetPasswordInsertCacheMut       sync.RWMutex
	resetPasswordInsertCache          = make(map[string]insertCache)
	resetPasswordUpdateCacheMut       sync.RWMutex
	resetPasswordUpdateCache          = make(map[string]updateCache)
	resetPasswordUpsertCacheMut       sync.RWMutex
	resetPasswordUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var resetPasswordAfterSelectHooks []ResetPasswordHook

var resetPasswordBeforeInsertHooks []ResetPasswordHook
var resetPasswordAfterInsertHooks []ResetPasswordHook

var resetPasswordBeforeUpdateHooks []ResetPasswordHook
var resetPasswordAfterUpdateHooks []ResetPasswordHook

var resetPasswordBeforeDeleteHooks []ResetPasswordHook
var resetPasswordAfterDeleteHooks []ResetPasswordHook

var resetPasswordBeforeUpsertHooks []ResetPasswordHook
var resetPasswordAfterUpsertHooks []ResetPasswordHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *ResetPassword) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range resetPasswordAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *ResetPassword) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range resetPasswordBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *ResetPassword) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range resetPasswordAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *ResetPassword) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range resetPasswordBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *ResetPassword) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range resetPasswordAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *ResetPassword) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range resetPasswordBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *ResetPassword) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range resetPasswordAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *ResetPassword) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range resetPasswordBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *ResetPassword) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range resetPasswordAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddResetPasswordHook registers your hook function for all future operations.
func AddResetPasswordHook(hookPoint boil.HookPoint, resetPasswordHook ResetPasswordHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		resetPasswordAfterSelectHooks = append(resetPasswordAfterSelectHooks, resetPasswordHook)
	case boil.BeforeInsertHook:
		resetPasswordBeforeInsertHooks = append(resetPasswordBeforeInsertHooks, resetPasswordHook)
	case boil.AfterInsertHook:
		resetPasswordAfterInsertHooks = append(resetPasswordAfterInsertHooks, resetPasswordHook)
	case boil.BeforeUpdateHook:
		resetPasswordBeforeUpdateHooks = append(resetPasswordBeforeUpdateHooks, resetPasswordHook)
	case boil.AfterUpdateHook:
		resetPasswordAfterUpdateHooks = append(resetPasswordAfterUpdateHooks, resetPasswordHook)
	case boil.BeforeDeleteHook:
		resetPasswordBeforeDeleteHooks = append(resetPasswordBeforeDeleteHooks, resetPasswordHook)
	case boil.AfterDeleteHook:
		resetPasswordAfterDeleteHooks = append(resetPasswordAfterDeleteHooks, resetPasswordHook)
	case boil.BeforeUpsertHook:
		resetPasswordBeforeUpsertHooks = append(resetPasswordBeforeUpsertHooks, resetPasswordHook)
	case boil.AfterUpsertHook:
		resetPasswordAfterUpsertHooks = append(resetPasswordAfterUpsertHooks, resetPasswordHook)
	}
}

// One returns a single resetPassword record from the query.
func (q resetPasswordQuery) One(exec boil.Executor) (*ResetPassword, error) {
	o := &ResetPassword{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(nil, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "boiler: failed to execute a one query for reset_passwords")
	}

	if err := o.doAfterSelectHooks(exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all ResetPassword records from the query.
func (q resetPasswordQuery) All(exec boil.Executor) (ResetPasswordSlice, error) {
	var o []*ResetPassword

	err := q.Bind(nil, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "boiler: failed to assign all query results to ResetPassword slice")
	}

	if len(resetPasswordAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all ResetPassword records in the query.
func (q resetPasswordQuery) Count(exec boil.Executor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: failed to count reset_passwords rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q resetPasswordQuery) Exists(exec boil.Executor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "boiler: failed to check if reset_passwords exists")
	}

	return count > 0, nil
}

// User pointed to by the foreign key.
func (o *ResetPassword) User(mods ...qm.QueryMod) userQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.UserID),
	}

	queryMods = append(queryMods, mods...)

	return Users(queryMods...)
}

// LoadUser allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (resetPasswordL) LoadUser(e boil.Executor, singular bool, maybeResetPassword interface{}, mods queries.Applicator) error {
	var slice []*ResetPassword
	var object *ResetPassword

	if singular {
		var ok bool
		object, ok = maybeResetPassword.(*ResetPassword)
		if !ok {
			object = new(ResetPassword)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeResetPassword)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeResetPassword))
			}
		}
	} else {
		s, ok := maybeResetPassword.(*[]*ResetPassword)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeResetPassword)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeResetPassword))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &resetPasswordR{}
		}
		args = append(args, object.UserID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &resetPasswordR{}
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
		foreign.R.ResetPassword = object
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.UserID == foreign.ID {
				local.R.User = foreign
				if foreign.R == nil {
					foreign.R = &userR{}
				}
				foreign.R.ResetPassword = local
				break
			}
		}
	}

	return nil
}

// SetUser of the resetPassword to the related item.
// Sets o.R.User to related.
// Adds o to related.R.ResetPassword.
func (o *ResetPassword) SetUser(exec boil.Executor, insert bool, related *User) error {
	var err error
	if insert {
		if err = related.Insert(exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"reset_passwords\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"user_id"}),
		strmangle.WhereClause("\"", "\"", 2, resetPasswordPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.UserID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}
	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.UserID = related.ID
	if o.R == nil {
		o.R = &resetPasswordR{
			User: related,
		}
	} else {
		o.R.User = related
	}

	if related.R == nil {
		related.R = &userR{
			ResetPassword: o,
		}
	} else {
		related.R.ResetPassword = o
	}

	return nil
}

// ResetPasswords retrieves all the records using an executor.
func ResetPasswords(mods ...qm.QueryMod) resetPasswordQuery {
	mods = append(mods, qm.From("\"reset_passwords\""), qmhelper.WhereIsNull("\"reset_passwords\".\"deleted_at\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"reset_passwords\".*"})
	}

	return resetPasswordQuery{q}
}

// FindResetPassword retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindResetPassword(exec boil.Executor, userID string, selectCols ...string) (*ResetPassword, error) {
	resetPasswordObj := &ResetPassword{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"reset_passwords\" where \"user_id\"=$1 and \"deleted_at\" is null", sel,
	)

	q := queries.Raw(query, userID)

	err := q.Bind(nil, exec, resetPasswordObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "boiler: unable to select from reset_passwords")
	}

	if err = resetPasswordObj.doAfterSelectHooks(exec); err != nil {
		return resetPasswordObj, err
	}

	return resetPasswordObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *ResetPassword) Insert(exec boil.Executor, columns boil.Columns) error {
	if o == nil {
		return errors.New("boiler: no reset_passwords provided for insertion")
	}

	var err error
	currTime := time.Now().In(boil.GetLocation())

	if o.UpdatedAt.IsZero() {
		o.UpdatedAt = currTime
	}
	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(resetPasswordColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	resetPasswordInsertCacheMut.RLock()
	cache, cached := resetPasswordInsertCache[key]
	resetPasswordInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			resetPasswordAllColumns,
			resetPasswordColumnsWithDefault,
			resetPasswordColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(resetPasswordType, resetPasswordMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(resetPasswordType, resetPasswordMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"reset_passwords\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"reset_passwords\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "boiler: unable to insert into reset_passwords")
	}

	if !cached {
		resetPasswordInsertCacheMut.Lock()
		resetPasswordInsertCache[key] = cache
		resetPasswordInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// Update uses an executor to update the ResetPassword.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *ResetPassword) Update(exec boil.Executor, columns boil.Columns) (int64, error) {
	currTime := time.Now().In(boil.GetLocation())

	o.UpdatedAt = currTime

	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	resetPasswordUpdateCacheMut.RLock()
	cache, cached := resetPasswordUpdateCache[key]
	resetPasswordUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			resetPasswordAllColumns,
			resetPasswordPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("boiler: unable to update reset_passwords, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"reset_passwords\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, resetPasswordPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(resetPasswordType, resetPasswordMapping, append(wl, resetPasswordPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "boiler: unable to update reset_passwords row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: failed to get rows affected by update for reset_passwords")
	}

	if !cached {
		resetPasswordUpdateCacheMut.Lock()
		resetPasswordUpdateCache[key] = cache
		resetPasswordUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(exec)
}

// UpdateAll updates all rows with the specified column values.
func (q resetPasswordQuery) UpdateAll(exec boil.Executor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to update all for reset_passwords")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to retrieve rows affected for reset_passwords")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ResetPasswordSlice) UpdateAll(exec boil.Executor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), resetPasswordPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"reset_passwords\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, resetPasswordPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to update all in resetPassword slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to retrieve rows affected all in update all resetPassword")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *ResetPassword) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("boiler: no reset_passwords provided for upsert")
	}
	currTime := time.Now().In(boil.GetLocation())

	o.UpdatedAt = currTime
	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(resetPasswordColumnsWithDefault, o)

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

	resetPasswordUpsertCacheMut.RLock()
	cache, cached := resetPasswordUpsertCache[key]
	resetPasswordUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			resetPasswordAllColumns,
			resetPasswordColumnsWithDefault,
			resetPasswordColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			resetPasswordAllColumns,
			resetPasswordPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("boiler: unable to upsert reset_passwords, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(resetPasswordPrimaryKeyColumns))
			copy(conflict, resetPasswordPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"reset_passwords\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(resetPasswordType, resetPasswordMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(resetPasswordType, resetPasswordMapping, ret)
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
		return errors.Wrap(err, "boiler: unable to upsert reset_passwords")
	}

	if !cached {
		resetPasswordUpsertCacheMut.Lock()
		resetPasswordUpsertCache[key] = cache
		resetPasswordUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// Delete deletes a single ResetPassword record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *ResetPassword) Delete(exec boil.Executor, hardDelete bool) (int64, error) {
	if o == nil {
		return 0, errors.New("boiler: no ResetPassword provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return 0, err
	}

	var (
		sql  string
		args []interface{}
	)
	if hardDelete {
		args = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), resetPasswordPrimaryKeyMapping)
		sql = "DELETE FROM \"reset_passwords\" WHERE \"user_id\"=$1"
	} else {
		currTime := time.Now().In(boil.GetLocation())
		o.DeletedAt = null.TimeFrom(currTime)
		wl := []string{"deleted_at"}
		sql = fmt.Sprintf("UPDATE \"reset_passwords\" SET %s WHERE \"user_id\"=$2",
			strmangle.SetParamNames("\"", "\"", 1, wl),
		)
		valueMapping, err := queries.BindMapping(resetPasswordType, resetPasswordMapping, append(wl, resetPasswordPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
		args = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), valueMapping)
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to delete from reset_passwords")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: failed to get rows affected by delete for reset_passwords")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q resetPasswordQuery) DeleteAll(exec boil.Executor, hardDelete bool) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("boiler: no resetPasswordQuery provided for delete all")
	}

	if hardDelete {
		queries.SetDelete(q.Query)
	} else {
		currTime := time.Now().In(boil.GetLocation())
		queries.SetUpdate(q.Query, M{"deleted_at": currTime})
	}

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to delete all from reset_passwords")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: failed to get rows affected by deleteall for reset_passwords")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ResetPasswordSlice) DeleteAll(exec boil.Executor, hardDelete bool) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(resetPasswordBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return 0, err
			}
		}
	}

	var (
		sql  string
		args []interface{}
	)
	if hardDelete {
		for _, obj := range o {
			pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), resetPasswordPrimaryKeyMapping)
			args = append(args, pkeyArgs...)
		}
		sql = "DELETE FROM \"reset_passwords\" WHERE " +
			strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, resetPasswordPrimaryKeyColumns, len(o))
	} else {
		currTime := time.Now().In(boil.GetLocation())
		for _, obj := range o {
			pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), resetPasswordPrimaryKeyMapping)
			args = append(args, pkeyArgs...)
			obj.DeletedAt = null.TimeFrom(currTime)
		}
		wl := []string{"deleted_at"}
		sql = fmt.Sprintf("UPDATE \"reset_passwords\" SET %s WHERE "+
			strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 2, resetPasswordPrimaryKeyColumns, len(o)),
			strmangle.SetParamNames("\"", "\"", 1, wl),
		)
		args = append([]interface{}{currTime}, args...)
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to delete all from resetPassword slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: failed to get rows affected by deleteall for reset_passwords")
	}

	if len(resetPasswordAfterDeleteHooks) != 0 {
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
func (o *ResetPassword) Reload(exec boil.Executor) error {
	ret, err := FindResetPassword(exec, o.UserID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ResetPasswordSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ResetPasswordSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), resetPasswordPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"reset_passwords\".* FROM \"reset_passwords\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, resetPasswordPrimaryKeyColumns, len(*o)) +
		"and \"deleted_at\" is null"

	q := queries.Raw(sql, args...)

	err := q.Bind(nil, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "boiler: unable to reload all in ResetPasswordSlice")
	}

	*o = slice

	return nil
}

// ResetPasswordExists checks if the ResetPassword row exists.
func ResetPasswordExists(exec boil.Executor, userID string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"reset_passwords\" where \"user_id\"=$1 and \"deleted_at\" is null limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, userID)
	}
	row := exec.QueryRow(sql, userID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "boiler: unable to check if reset_passwords exists")
	}

	return exists, nil
}

// Exists checks if the ResetPassword row exists.
func (o *ResetPassword) Exists(exec boil.Executor) (bool, error) {
	return ResetPasswordExists(exec, o.UserID)
}