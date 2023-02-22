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

// Privacy is an object representing the database table.
type Privacy struct {
	ID          string    `boiler:"id" boil:"id" json:"id" toml:"id" yaml:"id"`
	Name        string    `boiler:"name" boil:"name" json:"name" toml:"name" yaml:"name"`
	Slug        string    `boiler:"slug" boil:"slug" json:"slug" toml:"slug" yaml:"slug"`
	Description string    `boiler:"description" boil:"description" json:"description" toml:"description" yaml:"description"`
	CreatedAt   time.Time `boiler:"created_at" boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt   time.Time `boiler:"updated_at" boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`
	DeletedAt   null.Time `boiler:"deleted_at" boil:"deleted_at" json:"deleted_at,omitempty" toml:"deleted_at" yaml:"deleted_at,omitempty"`

	R *privacyR `boiler:"-" boil:"-" json:"-" toml:"-" yaml:"-"`
	L privacyL  `boiler:"-" boil:"-" json:"-" toml:"-" yaml:"-"`
}

var PrivacyColumns = struct {
	ID          string
	Name        string
	Slug        string
	Description string
	CreatedAt   string
	UpdatedAt   string
	DeletedAt   string
}{
	ID:          "id",
	Name:        "name",
	Slug:        "slug",
	Description: "description",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
	DeletedAt:   "deleted_at",
}

var PrivacyTableColumns = struct {
	ID          string
	Name        string
	Slug        string
	Description string
	CreatedAt   string
	UpdatedAt   string
	DeletedAt   string
}{
	ID:          "privacies.id",
	Name:        "privacies.name",
	Slug:        "privacies.slug",
	Description: "privacies.description",
	CreatedAt:   "privacies.created_at",
	UpdatedAt:   "privacies.updated_at",
	DeletedAt:   "privacies.deleted_at",
}

// Generated where

var PrivacyWhere = struct {
	ID          whereHelperstring
	Name        whereHelperstring
	Slug        whereHelperstring
	Description whereHelperstring
	CreatedAt   whereHelpertime_Time
	UpdatedAt   whereHelpertime_Time
	DeletedAt   whereHelpernull_Time
}{
	ID:          whereHelperstring{field: "\"privacies\".\"id\""},
	Name:        whereHelperstring{field: "\"privacies\".\"name\""},
	Slug:        whereHelperstring{field: "\"privacies\".\"slug\""},
	Description: whereHelperstring{field: "\"privacies\".\"description\""},
	CreatedAt:   whereHelpertime_Time{field: "\"privacies\".\"created_at\""},
	UpdatedAt:   whereHelpertime_Time{field: "\"privacies\".\"updated_at\""},
	DeletedAt:   whereHelpernull_Time{field: "\"privacies\".\"deleted_at\""},
}

// PrivacyRels is where relationship names are stored.
var PrivacyRels = struct {
	UserPrivacies string
}{
	UserPrivacies: "UserPrivacies",
}

// privacyR is where relationships are stored.
type privacyR struct {
	UserPrivacies UserPrivacySlice `boiler:"UserPrivacies" boil:"UserPrivacies" json:"UserPrivacies" toml:"UserPrivacies" yaml:"UserPrivacies"`
}

// NewStruct creates a new relationship struct
func (*privacyR) NewStruct() *privacyR {
	return &privacyR{}
}

func (r *privacyR) GetUserPrivacies() UserPrivacySlice {
	if r == nil {
		return nil
	}
	return r.UserPrivacies
}

// privacyL is where Load methods for each relationship are stored.
type privacyL struct{}

var (
	privacyAllColumns            = []string{"id", "name", "slug", "description", "created_at", "updated_at", "deleted_at"}
	privacyColumnsWithoutDefault = []string{"name", "slug", "description"}
	privacyColumnsWithDefault    = []string{"id", "created_at", "updated_at", "deleted_at"}
	privacyPrimaryKeyColumns     = []string{"id"}
	privacyGeneratedColumns      = []string{}
)

type (
	// PrivacySlice is an alias for a slice of pointers to Privacy.
	// This should almost always be used instead of []Privacy.
	PrivacySlice []*Privacy
	// PrivacyHook is the signature for custom Privacy hook methods
	PrivacyHook func(boil.Executor, *Privacy) error

	privacyQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	privacyType                 = reflect.TypeOf(&Privacy{})
	privacyMapping              = queries.MakeStructMapping(privacyType)
	privacyPrimaryKeyMapping, _ = queries.BindMapping(privacyType, privacyMapping, privacyPrimaryKeyColumns)
	privacyInsertCacheMut       sync.RWMutex
	privacyInsertCache          = make(map[string]insertCache)
	privacyUpdateCacheMut       sync.RWMutex
	privacyUpdateCache          = make(map[string]updateCache)
	privacyUpsertCacheMut       sync.RWMutex
	privacyUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var privacyAfterSelectHooks []PrivacyHook

var privacyBeforeInsertHooks []PrivacyHook
var privacyAfterInsertHooks []PrivacyHook

var privacyBeforeUpdateHooks []PrivacyHook
var privacyAfterUpdateHooks []PrivacyHook

var privacyBeforeDeleteHooks []PrivacyHook
var privacyAfterDeleteHooks []PrivacyHook

var privacyBeforeUpsertHooks []PrivacyHook
var privacyAfterUpsertHooks []PrivacyHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Privacy) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range privacyAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Privacy) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range privacyBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Privacy) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range privacyAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Privacy) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range privacyBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Privacy) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range privacyAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Privacy) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range privacyBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Privacy) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range privacyAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Privacy) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range privacyBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Privacy) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range privacyAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddPrivacyHook registers your hook function for all future operations.
func AddPrivacyHook(hookPoint boil.HookPoint, privacyHook PrivacyHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		privacyAfterSelectHooks = append(privacyAfterSelectHooks, privacyHook)
	case boil.BeforeInsertHook:
		privacyBeforeInsertHooks = append(privacyBeforeInsertHooks, privacyHook)
	case boil.AfterInsertHook:
		privacyAfterInsertHooks = append(privacyAfterInsertHooks, privacyHook)
	case boil.BeforeUpdateHook:
		privacyBeforeUpdateHooks = append(privacyBeforeUpdateHooks, privacyHook)
	case boil.AfterUpdateHook:
		privacyAfterUpdateHooks = append(privacyAfterUpdateHooks, privacyHook)
	case boil.BeforeDeleteHook:
		privacyBeforeDeleteHooks = append(privacyBeforeDeleteHooks, privacyHook)
	case boil.AfterDeleteHook:
		privacyAfterDeleteHooks = append(privacyAfterDeleteHooks, privacyHook)
	case boil.BeforeUpsertHook:
		privacyBeforeUpsertHooks = append(privacyBeforeUpsertHooks, privacyHook)
	case boil.AfterUpsertHook:
		privacyAfterUpsertHooks = append(privacyAfterUpsertHooks, privacyHook)
	}
}

// One returns a single privacy record from the query.
func (q privacyQuery) One(exec boil.Executor) (*Privacy, error) {
	o := &Privacy{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(nil, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "boiler: failed to execute a one query for privacies")
	}

	if err := o.doAfterSelectHooks(exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Privacy records from the query.
func (q privacyQuery) All(exec boil.Executor) (PrivacySlice, error) {
	var o []*Privacy

	err := q.Bind(nil, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "boiler: failed to assign all query results to Privacy slice")
	}

	if len(privacyAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Privacy records in the query.
func (q privacyQuery) Count(exec boil.Executor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: failed to count privacies rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q privacyQuery) Exists(exec boil.Executor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "boiler: failed to check if privacies exists")
	}

	return count > 0, nil
}

// UserPrivacies retrieves all the user_privacy's UserPrivacies with an executor.
func (o *Privacy) UserPrivacies(mods ...qm.QueryMod) userPrivacyQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"user_privacies\".\"privacy_id\"=?", o.ID),
	)

	return UserPrivacies(queryMods...)
}

// LoadUserPrivacies allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (privacyL) LoadUserPrivacies(e boil.Executor, singular bool, maybePrivacy interface{}, mods queries.Applicator) error {
	var slice []*Privacy
	var object *Privacy

	if singular {
		var ok bool
		object, ok = maybePrivacy.(*Privacy)
		if !ok {
			object = new(Privacy)
			ok = queries.SetFromEmbeddedStruct(&object, &maybePrivacy)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybePrivacy))
			}
		}
	} else {
		s, ok := maybePrivacy.(*[]*Privacy)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybePrivacy)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybePrivacy))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &privacyR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &privacyR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`user_privacies`),
		qm.WhereIn(`user_privacies.privacy_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.Query(e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load user_privacies")
	}

	var resultSlice []*UserPrivacy
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice user_privacies")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on user_privacies")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for user_privacies")
	}

	if len(userPrivacyAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.UserPrivacies = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &userPrivacyR{}
			}
			foreign.R.Privacy = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.PrivacyID {
				local.R.UserPrivacies = append(local.R.UserPrivacies, foreign)
				if foreign.R == nil {
					foreign.R = &userPrivacyR{}
				}
				foreign.R.Privacy = local
				break
			}
		}
	}

	return nil
}

// AddUserPrivacies adds the given related objects to the existing relationships
// of the privacy, optionally inserting them as new records.
// Appends related to o.R.UserPrivacies.
// Sets related.R.Privacy appropriately.
func (o *Privacy) AddUserPrivacies(exec boil.Executor, insert bool, related ...*UserPrivacy) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.PrivacyID = o.ID
			if err = rel.Insert(exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"user_privacies\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"privacy_id"}),
				strmangle.WhereClause("\"", "\"", 2, userPrivacyPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.UserID, rel.PrivacyID}

			if boil.DebugMode {
				fmt.Fprintln(boil.DebugWriter, updateQuery)
				fmt.Fprintln(boil.DebugWriter, values)
			}
			if _, err = exec.Exec(updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.PrivacyID = o.ID
		}
	}

	if o.R == nil {
		o.R = &privacyR{
			UserPrivacies: related,
		}
	} else {
		o.R.UserPrivacies = append(o.R.UserPrivacies, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &userPrivacyR{
				Privacy: o,
			}
		} else {
			rel.R.Privacy = o
		}
	}
	return nil
}

// Privacies retrieves all the records using an executor.
func Privacies(mods ...qm.QueryMod) privacyQuery {
	mods = append(mods, qm.From("\"privacies\""), qmhelper.WhereIsNull("\"privacies\".\"deleted_at\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"privacies\".*"})
	}

	return privacyQuery{q}
}

// FindPrivacy retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindPrivacy(exec boil.Executor, iD string, selectCols ...string) (*Privacy, error) {
	privacyObj := &Privacy{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"privacies\" where \"id\"=$1 and \"deleted_at\" is null", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(nil, exec, privacyObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "boiler: unable to select from privacies")
	}

	if err = privacyObj.doAfterSelectHooks(exec); err != nil {
		return privacyObj, err
	}

	return privacyObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Privacy) Insert(exec boil.Executor, columns boil.Columns) error {
	if o == nil {
		return errors.New("boiler: no privacies provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(privacyColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	privacyInsertCacheMut.RLock()
	cache, cached := privacyInsertCache[key]
	privacyInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			privacyAllColumns,
			privacyColumnsWithDefault,
			privacyColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(privacyType, privacyMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(privacyType, privacyMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"privacies\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"privacies\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "boiler: unable to insert into privacies")
	}

	if !cached {
		privacyInsertCacheMut.Lock()
		privacyInsertCache[key] = cache
		privacyInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// Update uses an executor to update the Privacy.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Privacy) Update(exec boil.Executor, columns boil.Columns) (int64, error) {
	currTime := time.Now().In(boil.GetLocation())

	o.UpdatedAt = currTime

	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	privacyUpdateCacheMut.RLock()
	cache, cached := privacyUpdateCache[key]
	privacyUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			privacyAllColumns,
			privacyPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("boiler: unable to update privacies, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"privacies\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, privacyPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(privacyType, privacyMapping, append(wl, privacyPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "boiler: unable to update privacies row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: failed to get rows affected by update for privacies")
	}

	if !cached {
		privacyUpdateCacheMut.Lock()
		privacyUpdateCache[key] = cache
		privacyUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(exec)
}

// UpdateAll updates all rows with the specified column values.
func (q privacyQuery) UpdateAll(exec boil.Executor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to update all for privacies")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to retrieve rows affected for privacies")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o PrivacySlice) UpdateAll(exec boil.Executor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), privacyPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"privacies\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, privacyPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to update all in privacy slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to retrieve rows affected all in update all privacy")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Privacy) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("boiler: no privacies provided for upsert")
	}
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}
	o.UpdatedAt = currTime

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(privacyColumnsWithDefault, o)

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

	privacyUpsertCacheMut.RLock()
	cache, cached := privacyUpsertCache[key]
	privacyUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			privacyAllColumns,
			privacyColumnsWithDefault,
			privacyColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			privacyAllColumns,
			privacyPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("boiler: unable to upsert privacies, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(privacyPrimaryKeyColumns))
			copy(conflict, privacyPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"privacies\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(privacyType, privacyMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(privacyType, privacyMapping, ret)
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
		return errors.Wrap(err, "boiler: unable to upsert privacies")
	}

	if !cached {
		privacyUpsertCacheMut.Lock()
		privacyUpsertCache[key] = cache
		privacyUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// Delete deletes a single Privacy record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Privacy) Delete(exec boil.Executor, hardDelete bool) (int64, error) {
	if o == nil {
		return 0, errors.New("boiler: no Privacy provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return 0, err
	}

	var (
		sql  string
		args []interface{}
	)
	if hardDelete {
		args = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), privacyPrimaryKeyMapping)
		sql = "DELETE FROM \"privacies\" WHERE \"id\"=$1"
	} else {
		currTime := time.Now().In(boil.GetLocation())
		o.DeletedAt = null.TimeFrom(currTime)
		wl := []string{"deleted_at"}
		sql = fmt.Sprintf("UPDATE \"privacies\" SET %s WHERE \"id\"=$2",
			strmangle.SetParamNames("\"", "\"", 1, wl),
		)
		valueMapping, err := queries.BindMapping(privacyType, privacyMapping, append(wl, privacyPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "boiler: unable to delete from privacies")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: failed to get rows affected by delete for privacies")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q privacyQuery) DeleteAll(exec boil.Executor, hardDelete bool) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("boiler: no privacyQuery provided for delete all")
	}

	if hardDelete {
		queries.SetDelete(q.Query)
	} else {
		currTime := time.Now().In(boil.GetLocation())
		queries.SetUpdate(q.Query, M{"deleted_at": currTime})
	}

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to delete all from privacies")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: failed to get rows affected by deleteall for privacies")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o PrivacySlice) DeleteAll(exec boil.Executor, hardDelete bool) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(privacyBeforeDeleteHooks) != 0 {
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
			pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), privacyPrimaryKeyMapping)
			args = append(args, pkeyArgs...)
		}
		sql = "DELETE FROM \"privacies\" WHERE " +
			strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, privacyPrimaryKeyColumns, len(o))
	} else {
		currTime := time.Now().In(boil.GetLocation())
		for _, obj := range o {
			pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), privacyPrimaryKeyMapping)
			args = append(args, pkeyArgs...)
			obj.DeletedAt = null.TimeFrom(currTime)
		}
		wl := []string{"deleted_at"}
		sql = fmt.Sprintf("UPDATE \"privacies\" SET %s WHERE "+
			strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 2, privacyPrimaryKeyColumns, len(o)),
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
		return 0, errors.Wrap(err, "boiler: unable to delete all from privacy slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: failed to get rows affected by deleteall for privacies")
	}

	if len(privacyAfterDeleteHooks) != 0 {
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
func (o *Privacy) Reload(exec boil.Executor) error {
	ret, err := FindPrivacy(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *PrivacySlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := PrivacySlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), privacyPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"privacies\".* FROM \"privacies\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, privacyPrimaryKeyColumns, len(*o)) +
		"and \"deleted_at\" is null"

	q := queries.Raw(sql, args...)

	err := q.Bind(nil, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "boiler: unable to reload all in PrivacySlice")
	}

	*o = slice

	return nil
}

// PrivacyExists checks if the Privacy row exists.
func PrivacyExists(exec boil.Executor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"privacies\" where \"id\"=$1 and \"deleted_at\" is null limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}
	row := exec.QueryRow(sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "boiler: unable to check if privacies exists")
	}

	return exists, nil
}

// Exists checks if the Privacy row exists.
func (o *Privacy) Exists(exec boil.Executor) (bool, error) {
	return PrivacyExists(exec, o.ID)
}