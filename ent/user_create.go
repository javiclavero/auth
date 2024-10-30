// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/javiclavero/go-auth-service/ent/user"
)

// UserCreate is the builder for creating a User entity.
type UserCreate struct {
	config
	mutation *UserMutation
	hooks    []Hook
}

// SetUseID sets the "use_id" field.
func (uc *UserCreate) SetUseID(u uint64) *UserCreate {
	uc.mutation.SetUseID(u)
	return uc
}

// SetUseType sets the "use_type" field.
func (uc *UserCreate) SetUseType(i int64) *UserCreate {
	uc.mutation.SetUseType(i)
	return uc
}

// SetNillableUseType sets the "use_type" field if the given value is not nil.
func (uc *UserCreate) SetNillableUseType(i *int64) *UserCreate {
	if i != nil {
		uc.SetUseType(*i)
	}
	return uc
}

// SetUseUsername sets the "use_username" field.
func (uc *UserCreate) SetUseUsername(s string) *UserCreate {
	uc.mutation.SetUseUsername(s)
	return uc
}

// SetUsePwd sets the "use_pwd" field.
func (uc *UserCreate) SetUsePwd(s string) *UserCreate {
	uc.mutation.SetUsePwd(s)
	return uc
}

// SetUseEmail sets the "use_email" field.
func (uc *UserCreate) SetUseEmail(s string) *UserCreate {
	uc.mutation.SetUseEmail(s)
	return uc
}

// SetUseResetToken sets the "use_reset_token" field.
func (uc *UserCreate) SetUseResetToken(s string) *UserCreate {
	uc.mutation.SetUseResetToken(s)
	return uc
}

// SetNillableUseResetToken sets the "use_reset_token" field if the given value is not nil.
func (uc *UserCreate) SetNillableUseResetToken(s *string) *UserCreate {
	if s != nil {
		uc.SetUseResetToken(*s)
	}
	return uc
}

// SetUseResetTokenExpiry sets the "use_reset_token_expiry" field.
func (uc *UserCreate) SetUseResetTokenExpiry(t time.Time) *UserCreate {
	uc.mutation.SetUseResetTokenExpiry(t)
	return uc
}

// SetNillableUseResetTokenExpiry sets the "use_reset_token_expiry" field if the given value is not nil.
func (uc *UserCreate) SetNillableUseResetTokenExpiry(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetUseResetTokenExpiry(*t)
	}
	return uc
}

// SetUseAccessToken sets the "use_access_token" field.
func (uc *UserCreate) SetUseAccessToken(s string) *UserCreate {
	uc.mutation.SetUseAccessToken(s)
	return uc
}

// SetNillableUseAccessToken sets the "use_access_token" field if the given value is not nil.
func (uc *UserCreate) SetNillableUseAccessToken(s *string) *UserCreate {
	if s != nil {
		uc.SetUseAccessToken(*s)
	}
	return uc
}

// SetUseAccessTokenExpiry sets the "use_access_token_expiry" field.
func (uc *UserCreate) SetUseAccessTokenExpiry(t time.Time) *UserCreate {
	uc.mutation.SetUseAccessTokenExpiry(t)
	return uc
}

// SetNillableUseAccessTokenExpiry sets the "use_access_token_expiry" field if the given value is not nil.
func (uc *UserCreate) SetNillableUseAccessTokenExpiry(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetUseAccessTokenExpiry(*t)
	}
	return uc
}

// SetUseCreatedAt sets the "use_created_at" field.
func (uc *UserCreate) SetUseCreatedAt(t time.Time) *UserCreate {
	uc.mutation.SetUseCreatedAt(t)
	return uc
}

// SetNillableUseCreatedAt sets the "use_created_at" field if the given value is not nil.
func (uc *UserCreate) SetNillableUseCreatedAt(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetUseCreatedAt(*t)
	}
	return uc
}

// SetUseUpdatedAt sets the "use_updated_at" field.
func (uc *UserCreate) SetUseUpdatedAt(t time.Time) *UserCreate {
	uc.mutation.SetUseUpdatedAt(t)
	return uc
}

// SetNillableUseUpdatedAt sets the "use_updated_at" field if the given value is not nil.
func (uc *UserCreate) SetNillableUseUpdatedAt(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetUseUpdatedAt(*t)
	}
	return uc
}

// SetUseDeletedAt sets the "use_deleted_at" field.
func (uc *UserCreate) SetUseDeletedAt(t time.Time) *UserCreate {
	uc.mutation.SetUseDeletedAt(t)
	return uc
}

// SetNillableUseDeletedAt sets the "use_deleted_at" field if the given value is not nil.
func (uc *UserCreate) SetNillableUseDeletedAt(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetUseDeletedAt(*t)
	}
	return uc
}

// Mutation returns the UserMutation object of the builder.
func (uc *UserCreate) Mutation() *UserMutation {
	return uc.mutation
}

// Save creates the User in the database.
func (uc *UserCreate) Save(ctx context.Context) (*User, error) {
	uc.defaults()
	return withHooks(ctx, uc.sqlSave, uc.mutation, uc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (uc *UserCreate) SaveX(ctx context.Context) *User {
	v, err := uc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uc *UserCreate) Exec(ctx context.Context) error {
	_, err := uc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uc *UserCreate) ExecX(ctx context.Context) {
	if err := uc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uc *UserCreate) defaults() {
	if _, ok := uc.mutation.UseType(); !ok {
		v := user.DefaultUseType
		uc.mutation.SetUseType(v)
	}
	if _, ok := uc.mutation.UseCreatedAt(); !ok {
		v := user.DefaultUseCreatedAt()
		uc.mutation.SetUseCreatedAt(v)
	}
	if _, ok := uc.mutation.UseUpdatedAt(); !ok {
		v := user.DefaultUseUpdatedAt()
		uc.mutation.SetUseUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uc *UserCreate) check() error {
	if _, ok := uc.mutation.UseID(); !ok {
		return &ValidationError{Name: "use_id", err: errors.New(`ent: missing required field "User.use_id"`)}
	}
	if _, ok := uc.mutation.UseType(); !ok {
		return &ValidationError{Name: "use_type", err: errors.New(`ent: missing required field "User.use_type"`)}
	}
	if _, ok := uc.mutation.UseUsername(); !ok {
		return &ValidationError{Name: "use_username", err: errors.New(`ent: missing required field "User.use_username"`)}
	}
	if v, ok := uc.mutation.UseUsername(); ok {
		if err := user.UseUsernameValidator(v); err != nil {
			return &ValidationError{Name: "use_username", err: fmt.Errorf(`ent: validator failed for field "User.use_username": %w`, err)}
		}
	}
	if _, ok := uc.mutation.UsePwd(); !ok {
		return &ValidationError{Name: "use_pwd", err: errors.New(`ent: missing required field "User.use_pwd"`)}
	}
	if v, ok := uc.mutation.UsePwd(); ok {
		if err := user.UsePwdValidator(v); err != nil {
			return &ValidationError{Name: "use_pwd", err: fmt.Errorf(`ent: validator failed for field "User.use_pwd": %w`, err)}
		}
	}
	if _, ok := uc.mutation.UseEmail(); !ok {
		return &ValidationError{Name: "use_email", err: errors.New(`ent: missing required field "User.use_email"`)}
	}
	if v, ok := uc.mutation.UseEmail(); ok {
		if err := user.UseEmailValidator(v); err != nil {
			return &ValidationError{Name: "use_email", err: fmt.Errorf(`ent: validator failed for field "User.use_email": %w`, err)}
		}
	}
	if v, ok := uc.mutation.UseResetToken(); ok {
		if err := user.UseResetTokenValidator(v); err != nil {
			return &ValidationError{Name: "use_reset_token", err: fmt.Errorf(`ent: validator failed for field "User.use_reset_token": %w`, err)}
		}
	}
	if v, ok := uc.mutation.UseAccessToken(); ok {
		if err := user.UseAccessTokenValidator(v); err != nil {
			return &ValidationError{Name: "use_access_token", err: fmt.Errorf(`ent: validator failed for field "User.use_access_token": %w`, err)}
		}
	}
	return nil
}

func (uc *UserCreate) sqlSave(ctx context.Context) (*User, error) {
	if err := uc.check(); err != nil {
		return nil, err
	}
	_node, _spec := uc.createSpec()
	if err := sqlgraph.CreateNode(ctx, uc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	uc.mutation.id = &_node.ID
	uc.mutation.done = true
	return _node, nil
}

func (uc *UserCreate) createSpec() (*User, *sqlgraph.CreateSpec) {
	var (
		_node = &User{config: uc.config}
		_spec = sqlgraph.NewCreateSpec(user.Table, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	)
	if value, ok := uc.mutation.UseID(); ok {
		_spec.SetField(user.FieldUseID, field.TypeUint64, value)
		_node.UseID = value
	}
	if value, ok := uc.mutation.UseType(); ok {
		_spec.SetField(user.FieldUseType, field.TypeInt64, value)
		_node.UseType = value
	}
	if value, ok := uc.mutation.UseUsername(); ok {
		_spec.SetField(user.FieldUseUsername, field.TypeString, value)
		_node.UseUsername = value
	}
	if value, ok := uc.mutation.UsePwd(); ok {
		_spec.SetField(user.FieldUsePwd, field.TypeString, value)
		_node.UsePwd = value
	}
	if value, ok := uc.mutation.UseEmail(); ok {
		_spec.SetField(user.FieldUseEmail, field.TypeString, value)
		_node.UseEmail = value
	}
	if value, ok := uc.mutation.UseResetToken(); ok {
		_spec.SetField(user.FieldUseResetToken, field.TypeString, value)
		_node.UseResetToken = &value
	}
	if value, ok := uc.mutation.UseResetTokenExpiry(); ok {
		_spec.SetField(user.FieldUseResetTokenExpiry, field.TypeTime, value)
		_node.UseResetTokenExpiry = &value
	}
	if value, ok := uc.mutation.UseAccessToken(); ok {
		_spec.SetField(user.FieldUseAccessToken, field.TypeString, value)
		_node.UseAccessToken = &value
	}
	if value, ok := uc.mutation.UseAccessTokenExpiry(); ok {
		_spec.SetField(user.FieldUseAccessTokenExpiry, field.TypeTime, value)
		_node.UseAccessTokenExpiry = &value
	}
	if value, ok := uc.mutation.UseCreatedAt(); ok {
		_spec.SetField(user.FieldUseCreatedAt, field.TypeTime, value)
		_node.UseCreatedAt = &value
	}
	if value, ok := uc.mutation.UseUpdatedAt(); ok {
		_spec.SetField(user.FieldUseUpdatedAt, field.TypeTime, value)
		_node.UseUpdatedAt = &value
	}
	if value, ok := uc.mutation.UseDeletedAt(); ok {
		_spec.SetField(user.FieldUseDeletedAt, field.TypeTime, value)
		_node.UseDeletedAt = &value
	}
	return _node, _spec
}

// UserCreateBulk is the builder for creating many User entities in bulk.
type UserCreateBulk struct {
	config
	err      error
	builders []*UserCreate
}

// Save creates the User entities in the database.
func (ucb *UserCreateBulk) Save(ctx context.Context) ([]*User, error) {
	if ucb.err != nil {
		return nil, ucb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ucb.builders))
	nodes := make([]*User, len(ucb.builders))
	mutators := make([]Mutator, len(ucb.builders))
	for i := range ucb.builders {
		func(i int, root context.Context) {
			builder := ucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ucb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ucb *UserCreateBulk) SaveX(ctx context.Context) []*User {
	v, err := ucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ucb *UserCreateBulk) Exec(ctx context.Context) error {
	_, err := ucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucb *UserCreateBulk) ExecX(ctx context.Context) {
	if err := ucb.Exec(ctx); err != nil {
		panic(err)
	}
}
