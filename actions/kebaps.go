package actions

import (

  "fmt"
  "net/http"
  "github.com/gobuffalo/buffalo"
  "github.com/gobuffalo/pop"
  "github.com/h3rby7/kebap/models"
)

// This file is generated by Buffalo. It offers a basic structure for
// adding, editing and deleting a page. If your model is more
// complex or you need more than the basic implementation you need to
// edit this file.

// Following naming logic is implemented in Buffalo:
// Model: Singular (Kebap)
// DB Table: Plural (kebaps)
// Resource: Plural (Kebaps)
// Path: Plural (/kebaps)
// View Template Folder: Plural (/templates/kebaps/)

// KebapsResource is the resource for the Kebap model
type KebapsResource struct{
  buffalo.Resource
}

// List gets all Kebaps. This function is mapped to the path
// GET /kebaps
func (v KebapsResource) List(c buffalo.Context) error {
  // Get the DB connection from the context
  tx, ok := c.Value("tx").(*pop.Connection)
  if !ok {
    return fmt.Errorf("no transaction found")
  }

  kebaps := &models.Kebaps{}

  // Paginate results. Params "page" and "per_page" control pagination.
  // Default values are "page=1" and "per_page=20".
  q := tx.PaginateFromParams(c.Params())

  // Retrieve all Kebaps from the DB
  if err := q.All(kebaps); err != nil {
    return err
  }

  // Add the paginator to the context so it can be used in the template.
  c.Set("pagination", q.Paginator)

  return c.Render(http.StatusOK, r.Auto(c, kebaps))
}

// Show gets the data for one Kebap. This function is mapped to
// the path GET /kebaps/{kebap_id}
func (v KebapsResource) Show(c buffalo.Context) error {
  // Get the DB connection from the context
  tx, ok := c.Value("tx").(*pop.Connection)
  if !ok {
    return fmt.Errorf("no transaction found")
  }

  // Allocate an empty Kebap
  kebap := &models.Kebap{}

  // To find the Kebap the parameter kebap_id is used.
  if err := tx.Find(kebap, c.Param("kebap_id")); err != nil {
    return c.Error(http.StatusNotFound, err)
  }

  return c.Render(http.StatusOK, r.Auto(c, kebap))
}

// New renders the form for creating a new Kebap.
// This function is mapped to the path GET /kebaps/new
func (v KebapsResource) New(c buffalo.Context) error {
  return c.Render(http.StatusOK, r.Auto(c, &models.Kebap{}))
}
// Create adds a Kebap to the DB. This function is mapped to the
// path POST /kebaps
func (v KebapsResource) Create(c buffalo.Context) error {
  // Allocate an empty Kebap
  kebap := &models.Kebap{}

  // Bind kebap to the html form elements
  if err := c.Bind(kebap); err != nil {
    return err
  }

  // Get the DB connection from the context
  tx, ok := c.Value("tx").(*pop.Connection)
  if !ok {
    return fmt.Errorf("no transaction found")
  }

  // Validate the data from the html form
  verrs, err := tx.ValidateAndCreate(kebap)
  if err != nil {
    return err
  }

  if verrs.HasAny() {
    // Make the errors available inside the html template
    c.Set("errors", verrs)

    // Render again the new.html template that the user can
    // correct the input.
    return c.Render(http.StatusUnprocessableEntity, r.Auto(c, kebap))
  }

  // If there are no errors set a success message
  c.Flash().Add("success", T.Translate(c, "kebap.created.success"))
  // and redirect to the kebaps index page
  return c.Render(http.StatusCreated, r.Auto(c, kebap))
}

// Edit renders a edit form for a Kebap. This function is
// mapped to the path GET /kebaps/{kebap_id}/edit
func (v KebapsResource) Edit(c buffalo.Context) error {
  // Get the DB connection from the context
  tx, ok := c.Value("tx").(*pop.Connection)
  if !ok {
    return fmt.Errorf("no transaction found")
  }

  // Allocate an empty Kebap
  kebap := &models.Kebap{}

  if err := tx.Find(kebap, c.Param("kebap_id")); err != nil {
    return c.Error(http.StatusNotFound, err)
  }

  return c.Render(http.StatusOK, r.Auto(c, kebap))
}
// Update changes a Kebap in the DB. This function is mapped to
// the path PUT /kebaps/{kebap_id}
func (v KebapsResource) Update(c buffalo.Context) error {
  // Get the DB connection from the context
  tx, ok := c.Value("tx").(*pop.Connection)
  if !ok {
    return fmt.Errorf("no transaction found")
  }

  // Allocate an empty Kebap
  kebap := &models.Kebap{}

  if err := tx.Find(kebap, c.Param("kebap_id")); err != nil {
    return c.Error(http.StatusNotFound, err)
  }

  // Bind Kebap to the html form elements
  if err := c.Bind(kebap); err != nil {
    return err
  }

  verrs, err := tx.ValidateAndUpdate(kebap)
  if err != nil {
    return err
  }

  if verrs.HasAny() {
    // Make the errors available inside the html template
    c.Set("errors", verrs)

    // Render again the edit.html template that the user can
    // correct the input.
    return c.Render(http.StatusUnprocessableEntity, r.Auto(c, kebap))
  }

  // If there are no errors set a success message
  c.Flash().Add("success", T.Translate(c, "kebap.updated.success"))
  // and redirect to the kebaps index page
  return c.Render(http.StatusOK, r.Auto(c, kebap))
}

// Destroy deletes a Kebap from the DB. This function is mapped
// to the path DELETE /kebaps/{kebap_id}
func (v KebapsResource) Destroy(c buffalo.Context) error {
  // Get the DB connection from the context
  tx, ok := c.Value("tx").(*pop.Connection)
  if !ok {
    return fmt.Errorf("no transaction found")
  }

  // Allocate an empty Kebap
  kebap := &models.Kebap{}

  // To find the Kebap the parameter kebap_id is used.
  if err := tx.Find(kebap, c.Param("kebap_id")); err != nil {
    return c.Error(http.StatusNotFound, err)
  }

  if err := tx.Destroy(kebap); err != nil {
    return err
  }

  // If there are no errors set a flash message
  c.Flash().Add("success", T.Translate(c, "kebap.destroyed.success"))
  // Redirect to the kebaps index page
  return c.Render(http.StatusOK, r.Auto(c, kebap))
}