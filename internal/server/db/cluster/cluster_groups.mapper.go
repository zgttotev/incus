//go:build linux && cgo && !agent

// Code generated by generate-database from the incus project - DO NOT EDIT.

package cluster

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
)

var clusterGroupObjects = RegisterStmt(`
SELECT cluster_groups.id, cluster_groups.name, coalesce(cluster_groups.description, '')
  FROM cluster_groups
  ORDER BY cluster_groups.name
`)

var clusterGroupObjectsByName = RegisterStmt(`
SELECT cluster_groups.id, cluster_groups.name, coalesce(cluster_groups.description, '')
  FROM cluster_groups
  WHERE ( cluster_groups.name = ? )
  ORDER BY cluster_groups.name
`)

var clusterGroupID = RegisterStmt(`
SELECT cluster_groups.id FROM cluster_groups
  WHERE cluster_groups.name = ?
`)

var clusterGroupCreate = RegisterStmt(`
INSERT INTO cluster_groups (name, description)
  VALUES (?, ?)
`)

var clusterGroupRename = RegisterStmt(`
UPDATE cluster_groups SET name = ? WHERE name = ?
`)

var clusterGroupDeleteByName = RegisterStmt(`
DELETE FROM cluster_groups WHERE name = ?
`)

var clusterGroupUpdate = RegisterStmt(`
UPDATE cluster_groups
  SET name = ?, description = ?
 WHERE id = ?
`)

// clusterGroupColumns returns a string of column names to be used with a SELECT statement for the entity.
// Use this function when building statements to retrieve database entries matching the ClusterGroup entity.
func clusterGroupColumns() string {
	return "clusters_groups.id, clusters_groups.name, coalesce(clusters_groups.description, '')"
}

// getClusterGroups can be used to run handwritten sql.Stmts to return a slice of objects.
func getClusterGroups(ctx context.Context, stmt *sql.Stmt, args ...any) ([]ClusterGroup, error) {
	objects := make([]ClusterGroup, 0)

	dest := func(scan func(dest ...any) error) error {
		c := ClusterGroup{}
		err := scan(&c.ID, &c.Name, &c.Description)
		if err != nil {
			return err
		}

		objects = append(objects, c)

		return nil
	}

	err := selectObjects(ctx, stmt, dest, args...)
	if err != nil {
		return nil, fmt.Errorf("Failed to fetch from \"clusters_groups\" table: %w", err)
	}

	return objects, nil
}

// getClusterGroupsRaw can be used to run handwritten query strings to return a slice of objects.
func getClusterGroupsRaw(ctx context.Context, db dbtx, sql string, args ...any) ([]ClusterGroup, error) {
	objects := make([]ClusterGroup, 0)

	dest := func(scan func(dest ...any) error) error {
		c := ClusterGroup{}
		err := scan(&c.ID, &c.Name, &c.Description)
		if err != nil {
			return err
		}

		objects = append(objects, c)

		return nil
	}

	err := scan(ctx, db, sql, dest, args...)
	if err != nil {
		return nil, fmt.Errorf("Failed to fetch from \"clusters_groups\" table: %w", err)
	}

	return objects, nil
}

// GetClusterGroups returns all available cluster_groups.
// generator: cluster_group GetMany
func GetClusterGroups(ctx context.Context, db dbtx, filters ...ClusterGroupFilter) (_ []ClusterGroup, _err error) {
	defer func() {
		_err = mapErr(_err, "Cluster_group")
	}()

	var err error

	// Result slice.
	objects := make([]ClusterGroup, 0)

	// Pick the prepared statement and arguments to use based on active criteria.
	var sqlStmt *sql.Stmt
	args := []any{}
	queryParts := [2]string{}

	if len(filters) == 0 {
		sqlStmt, err = Stmt(db, clusterGroupObjects)
		if err != nil {
			return nil, fmt.Errorf("Failed to get \"clusterGroupObjects\" prepared statement: %w", err)
		}
	}

	for i, filter := range filters {
		if filter.Name != nil && filter.ID == nil {
			args = append(args, []any{filter.Name}...)
			if len(filters) == 1 {
				sqlStmt, err = Stmt(db, clusterGroupObjectsByName)
				if err != nil {
					return nil, fmt.Errorf("Failed to get \"clusterGroupObjectsByName\" prepared statement: %w", err)
				}

				break
			}

			query, err := StmtString(clusterGroupObjectsByName)
			if err != nil {
				return nil, fmt.Errorf("Failed to get \"clusterGroupObjects\" prepared statement: %w", err)
			}

			parts := strings.SplitN(query, "ORDER BY", 2)
			if i == 0 {
				copy(queryParts[:], parts)
				continue
			}

			_, where, _ := strings.Cut(parts[0], "WHERE")
			queryParts[0] += "OR" + where
		} else if filter.ID == nil && filter.Name == nil {
			return nil, fmt.Errorf("Cannot filter on empty ClusterGroupFilter")
		} else {
			return nil, fmt.Errorf("No statement exists for the given Filter")
		}
	}

	// Select.
	if sqlStmt != nil {
		objects, err = getClusterGroups(ctx, sqlStmt, args...)
	} else {
		queryStr := strings.Join(queryParts[:], "ORDER BY")
		objects, err = getClusterGroupsRaw(ctx, db, queryStr, args...)
	}

	if err != nil {
		return nil, fmt.Errorf("Failed to fetch from \"clusters_groups\" table: %w", err)
	}

	return objects, nil
}

// GetClusterGroup returns the cluster_group with the given key.
// generator: cluster_group GetOne
func GetClusterGroup(ctx context.Context, db dbtx, name string) (_ *ClusterGroup, _err error) {
	defer func() {
		_err = mapErr(_err, "Cluster_group")
	}()

	filter := ClusterGroupFilter{}
	filter.Name = &name

	objects, err := GetClusterGroups(ctx, db, filter)
	if err != nil {
		return nil, fmt.Errorf("Failed to fetch from \"clusters_groups\" table: %w", err)
	}

	switch len(objects) {
	case 0:
		return nil, ErrNotFound
	case 1:
		return &objects[0], nil
	default:
		return nil, fmt.Errorf("More than one \"clusters_groups\" entry matches")
	}
}

// GetClusterGroupID return the ID of the cluster_group with the given key.
// generator: cluster_group ID
func GetClusterGroupID(ctx context.Context, db dbtx, name string) (_ int64, _err error) {
	defer func() {
		_err = mapErr(_err, "Cluster_group")
	}()

	stmt, err := Stmt(db, clusterGroupID)
	if err != nil {
		return -1, fmt.Errorf("Failed to get \"clusterGroupID\" prepared statement: %w", err)
	}

	row := stmt.QueryRowContext(ctx, name)
	var id int64
	err = row.Scan(&id)
	if errors.Is(err, sql.ErrNoRows) {
		return -1, ErrNotFound
	}

	if err != nil {
		return -1, fmt.Errorf("Failed to get \"clusters_groups\" ID: %w", err)
	}

	return id, nil
}

// ClusterGroupExists checks if a cluster_group with the given key exists.
// generator: cluster_group Exists
func ClusterGroupExists(ctx context.Context, db dbtx, name string) (_ bool, _err error) {
	defer func() {
		_err = mapErr(_err, "Cluster_group")
	}()

	stmt, err := Stmt(db, clusterGroupID)
	if err != nil {
		return false, fmt.Errorf("Failed to get \"clusterGroupID\" prepared statement: %w", err)
	}

	row := stmt.QueryRowContext(ctx, name)
	var id int64
	err = row.Scan(&id)
	if errors.Is(err, sql.ErrNoRows) {
		return false, nil
	}

	if err != nil {
		return false, fmt.Errorf("Failed to get \"clusters_groups\" ID: %w", err)
	}

	return true, nil
}

// RenameClusterGroup renames the cluster_group matching the given key parameters.
// generator: cluster_group Rename
func RenameClusterGroup(ctx context.Context, db dbtx, name string, to string) (_err error) {
	defer func() {
		_err = mapErr(_err, "Cluster_group")
	}()

	stmt, err := Stmt(db, clusterGroupRename)
	if err != nil {
		return fmt.Errorf("Failed to get \"clusterGroupRename\" prepared statement: %w", err)
	}

	result, err := stmt.Exec(to, name)
	if err != nil {
		return fmt.Errorf("Rename ClusterGroup failed: %w", err)
	}

	n, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("Fetch affected rows failed: %w", err)
	}

	if n != 1 {
		return fmt.Errorf("Query affected %d rows instead of 1", n)
	}

	return nil
}

// CreateClusterGroup adds a new cluster_group to the database.
// generator: cluster_group Create
func CreateClusterGroup(ctx context.Context, db dbtx, object ClusterGroup) (_ int64, _err error) {
	defer func() {
		_err = mapErr(_err, "Cluster_group")
	}()

	// Check if a cluster_group with the same key exists.
	exists, err := ClusterGroupExists(ctx, db, object.Name)
	if err != nil {
		return -1, fmt.Errorf("Failed to check for duplicates: %w", err)
	}

	if exists {
		return -1, ErrConflict
	}

	args := make([]any, 2)

	// Populate the statement arguments.
	args[0] = object.Name
	args[1] = object.Description

	// Prepared statement to use.
	stmt, err := Stmt(db, clusterGroupCreate)
	if err != nil {
		return -1, fmt.Errorf("Failed to get \"clusterGroupCreate\" prepared statement: %w", err)
	}

	// Execute the statement.
	result, err := stmt.Exec(args...)
	if err != nil {
		return -1, fmt.Errorf("Failed to create \"clusters_groups\" entry: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return -1, fmt.Errorf("Failed to fetch \"clusters_groups\" entry ID: %w", err)
	}

	return id, nil
}

// UpdateClusterGroup updates the cluster_group matching the given key parameters.
// generator: cluster_group Update
func UpdateClusterGroup(ctx context.Context, db dbtx, name string, object ClusterGroup) (_err error) {
	defer func() {
		_err = mapErr(_err, "Cluster_group")
	}()

	id, err := GetClusterGroupID(ctx, db, name)
	if err != nil {
		return err
	}

	stmt, err := Stmt(db, clusterGroupUpdate)
	if err != nil {
		return fmt.Errorf("Failed to get \"clusterGroupUpdate\" prepared statement: %w", err)
	}

	result, err := stmt.Exec(object.Name, object.Description, id)
	if err != nil {
		return fmt.Errorf("Update \"clusters_groups\" entry failed: %w", err)
	}

	n, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("Fetch affected rows: %w", err)
	}

	if n != 1 {
		return fmt.Errorf("Query updated %d rows instead of 1", n)
	}

	return nil
}

// DeleteClusterGroup deletes the cluster_group matching the given key parameters.
// generator: cluster_group DeleteOne-by-Name
func DeleteClusterGroup(ctx context.Context, db dbtx, name string) (_err error) {
	defer func() {
		_err = mapErr(_err, "Cluster_group")
	}()

	stmt, err := Stmt(db, clusterGroupDeleteByName)
	if err != nil {
		return fmt.Errorf("Failed to get \"clusterGroupDeleteByName\" prepared statement: %w", err)
	}

	result, err := stmt.Exec(name)
	if err != nil {
		return fmt.Errorf("Delete \"clusters_groups\": %w", err)
	}

	n, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("Fetch affected rows: %w", err)
	}

	if n == 0 {
		return ErrNotFound
	} else if n > 1 {
		return fmt.Errorf("Query deleted %d ClusterGroup rows instead of 1", n)
	}

	return nil
}
