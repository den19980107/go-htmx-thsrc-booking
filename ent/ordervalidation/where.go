// Code generated by ent, DO NOT EDIT.

package ordervalidation

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/mikestefanello/pagoda/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.OrderValidation {
	return predicate.OrderValidation(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.OrderValidation {
	return predicate.OrderValidation(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.OrderValidation {
	return predicate.OrderValidation(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.OrderValidation {
	return predicate.OrderValidation(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.OrderValidation {
	return predicate.OrderValidation(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.OrderValidation {
	return predicate.OrderValidation(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.OrderValidation {
	return predicate.OrderValidation(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.OrderValidation {
	return predicate.OrderValidation(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.OrderValidation {
	return predicate.OrderValidation(sql.FieldLTE(FieldID, id))
}

// JessionID applies equality check predicate on the "jession_id" field. It's identical to JessionIDEQ.
func JessionID(v string) predicate.OrderValidation {
	return predicate.OrderValidation(sql.FieldEQ(FieldJessionID, v))
}

// CaptchaImage applies equality check predicate on the "captcha_image" field. It's identical to CaptchaImageEQ.
func CaptchaImage(v string) predicate.OrderValidation {
	return predicate.OrderValidation(sql.FieldEQ(FieldCaptchaImage, v))
}

// Cookies applies equality check predicate on the "cookies" field. It's identical to CookiesEQ.
func Cookies(v string) predicate.OrderValidation {
	return predicate.OrderValidation(sql.FieldEQ(FieldCookies, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.OrderValidation {
	return predicate.OrderValidation(sql.FieldEQ(FieldCreatedAt, v))
}

// JessionIDEQ applies the EQ predicate on the "jession_id" field.
func JessionIDEQ(v string) predicate.OrderValidation {
	return predicate.OrderValidation(sql.FieldEQ(FieldJessionID, v))
}

// JessionIDNEQ applies the NEQ predicate on the "jession_id" field.
func JessionIDNEQ(v string) predicate.OrderValidation {
	return predicate.OrderValidation(sql.FieldNEQ(FieldJessionID, v))
}

// JessionIDIn applies the In predicate on the "jession_id" field.
func JessionIDIn(vs ...string) predicate.OrderValidation {
	return predicate.OrderValidation(sql.FieldIn(FieldJessionID, vs...))
}

// JessionIDNotIn applies the NotIn predicate on the "jession_id" field.
func JessionIDNotIn(vs ...string) predicate.OrderValidation {
	return predicate.OrderValidation(sql.FieldNotIn(FieldJessionID, vs...))
}

// JessionIDGT applies the GT predicate on the "jession_id" field.
func JessionIDGT(v string) predicate.OrderValidation {
	return predicate.OrderValidation(sql.FieldGT(FieldJessionID, v))
}

// JessionIDGTE applies the GTE predicate on the "jession_id" field.
func JessionIDGTE(v string) predicate.OrderValidation {
	return predicate.OrderValidation(sql.FieldGTE(FieldJessionID, v))
}

// JessionIDLT applies the LT predicate on the "jession_id" field.
func JessionIDLT(v string) predicate.OrderValidation {
	return predicate.OrderValidation(sql.FieldLT(FieldJessionID, v))
}

// JessionIDLTE applies the LTE predicate on the "jession_id" field.
func JessionIDLTE(v string) predicate.OrderValidation {
	return predicate.OrderValidation(sql.FieldLTE(FieldJessionID, v))
}

// JessionIDContains applies the Contains predicate on the "jession_id" field.
func JessionIDContains(v string) predicate.OrderValidation {
	return predicate.OrderValidation(sql.FieldContains(FieldJessionID, v))
}

// JessionIDHasPrefix applies the HasPrefix predicate on the "jession_id" field.
func JessionIDHasPrefix(v string) predicate.OrderValidation {
	return predicate.OrderValidation(sql.FieldHasPrefix(FieldJessionID, v))
}

// JessionIDHasSuffix applies the HasSuffix predicate on the "jession_id" field.
func JessionIDHasSuffix(v string) predicate.OrderValidation {
	return predicate.OrderValidation(sql.FieldHasSuffix(FieldJessionID, v))
}

// JessionIDEqualFold applies the EqualFold predicate on the "jession_id" field.
func JessionIDEqualFold(v string) predicate.OrderValidation {
	return predicate.OrderValidation(sql.FieldEqualFold(FieldJessionID, v))
}

// JessionIDContainsFold applies the ContainsFold predicate on the "jession_id" field.
func JessionIDContainsFold(v string) predicate.OrderValidation {
	return predicate.OrderValidation(sql.FieldContainsFold(FieldJessionID, v))
}

// CaptchaImageEQ applies the EQ predicate on the "captcha_image" field.
func CaptchaImageEQ(v string) predicate.OrderValidation {
	return predicate.OrderValidation(sql.FieldEQ(FieldCaptchaImage, v))
}

// CaptchaImageNEQ applies the NEQ predicate on the "captcha_image" field.
func CaptchaImageNEQ(v string) predicate.OrderValidation {
	return predicate.OrderValidation(sql.FieldNEQ(FieldCaptchaImage, v))
}

// CaptchaImageIn applies the In predicate on the "captcha_image" field.
func CaptchaImageIn(vs ...string) predicate.OrderValidation {
	return predicate.OrderValidation(sql.FieldIn(FieldCaptchaImage, vs...))
}

// CaptchaImageNotIn applies the NotIn predicate on the "captcha_image" field.
func CaptchaImageNotIn(vs ...string) predicate.OrderValidation {
	return predicate.OrderValidation(sql.FieldNotIn(FieldCaptchaImage, vs...))
}

// CaptchaImageGT applies the GT predicate on the "captcha_image" field.
func CaptchaImageGT(v string) predicate.OrderValidation {
	return predicate.OrderValidation(sql.FieldGT(FieldCaptchaImage, v))
}

// CaptchaImageGTE applies the GTE predicate on the "captcha_image" field.
func CaptchaImageGTE(v string) predicate.OrderValidation {
	return predicate.OrderValidation(sql.FieldGTE(FieldCaptchaImage, v))
}

// CaptchaImageLT applies the LT predicate on the "captcha_image" field.
func CaptchaImageLT(v string) predicate.OrderValidation {
	return predicate.OrderValidation(sql.FieldLT(FieldCaptchaImage, v))
}

// CaptchaImageLTE applies the LTE predicate on the "captcha_image" field.
func CaptchaImageLTE(v string) predicate.OrderValidation {
	return predicate.OrderValidation(sql.FieldLTE(FieldCaptchaImage, v))
}

// CaptchaImageContains applies the Contains predicate on the "captcha_image" field.
func CaptchaImageContains(v string) predicate.OrderValidation {
	return predicate.OrderValidation(sql.FieldContains(FieldCaptchaImage, v))
}

// CaptchaImageHasPrefix applies the HasPrefix predicate on the "captcha_image" field.
func CaptchaImageHasPrefix(v string) predicate.OrderValidation {
	return predicate.OrderValidation(sql.FieldHasPrefix(FieldCaptchaImage, v))
}

// CaptchaImageHasSuffix applies the HasSuffix predicate on the "captcha_image" field.
func CaptchaImageHasSuffix(v string) predicate.OrderValidation {
	return predicate.OrderValidation(sql.FieldHasSuffix(FieldCaptchaImage, v))
}

// CaptchaImageEqualFold applies the EqualFold predicate on the "captcha_image" field.
func CaptchaImageEqualFold(v string) predicate.OrderValidation {
	return predicate.OrderValidation(sql.FieldEqualFold(FieldCaptchaImage, v))
}

// CaptchaImageContainsFold applies the ContainsFold predicate on the "captcha_image" field.
func CaptchaImageContainsFold(v string) predicate.OrderValidation {
	return predicate.OrderValidation(sql.FieldContainsFold(FieldCaptchaImage, v))
}

// CookiesEQ applies the EQ predicate on the "cookies" field.
func CookiesEQ(v string) predicate.OrderValidation {
	return predicate.OrderValidation(sql.FieldEQ(FieldCookies, v))
}

// CookiesNEQ applies the NEQ predicate on the "cookies" field.
func CookiesNEQ(v string) predicate.OrderValidation {
	return predicate.OrderValidation(sql.FieldNEQ(FieldCookies, v))
}

// CookiesIn applies the In predicate on the "cookies" field.
func CookiesIn(vs ...string) predicate.OrderValidation {
	return predicate.OrderValidation(sql.FieldIn(FieldCookies, vs...))
}

// CookiesNotIn applies the NotIn predicate on the "cookies" field.
func CookiesNotIn(vs ...string) predicate.OrderValidation {
	return predicate.OrderValidation(sql.FieldNotIn(FieldCookies, vs...))
}

// CookiesGT applies the GT predicate on the "cookies" field.
func CookiesGT(v string) predicate.OrderValidation {
	return predicate.OrderValidation(sql.FieldGT(FieldCookies, v))
}

// CookiesGTE applies the GTE predicate on the "cookies" field.
func CookiesGTE(v string) predicate.OrderValidation {
	return predicate.OrderValidation(sql.FieldGTE(FieldCookies, v))
}

// CookiesLT applies the LT predicate on the "cookies" field.
func CookiesLT(v string) predicate.OrderValidation {
	return predicate.OrderValidation(sql.FieldLT(FieldCookies, v))
}

// CookiesLTE applies the LTE predicate on the "cookies" field.
func CookiesLTE(v string) predicate.OrderValidation {
	return predicate.OrderValidation(sql.FieldLTE(FieldCookies, v))
}

// CookiesContains applies the Contains predicate on the "cookies" field.
func CookiesContains(v string) predicate.OrderValidation {
	return predicate.OrderValidation(sql.FieldContains(FieldCookies, v))
}

// CookiesHasPrefix applies the HasPrefix predicate on the "cookies" field.
func CookiesHasPrefix(v string) predicate.OrderValidation {
	return predicate.OrderValidation(sql.FieldHasPrefix(FieldCookies, v))
}

// CookiesHasSuffix applies the HasSuffix predicate on the "cookies" field.
func CookiesHasSuffix(v string) predicate.OrderValidation {
	return predicate.OrderValidation(sql.FieldHasSuffix(FieldCookies, v))
}

// CookiesEqualFold applies the EqualFold predicate on the "cookies" field.
func CookiesEqualFold(v string) predicate.OrderValidation {
	return predicate.OrderValidation(sql.FieldEqualFold(FieldCookies, v))
}

// CookiesContainsFold applies the ContainsFold predicate on the "cookies" field.
func CookiesContainsFold(v string) predicate.OrderValidation {
	return predicate.OrderValidation(sql.FieldContainsFold(FieldCookies, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.OrderValidation {
	return predicate.OrderValidation(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.OrderValidation {
	return predicate.OrderValidation(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.OrderValidation {
	return predicate.OrderValidation(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.OrderValidation {
	return predicate.OrderValidation(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.OrderValidation {
	return predicate.OrderValidation(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.OrderValidation {
	return predicate.OrderValidation(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.OrderValidation {
	return predicate.OrderValidation(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.OrderValidation {
	return predicate.OrderValidation(sql.FieldLTE(FieldCreatedAt, v))
}

// HasOrder applies the HasEdge predicate on the "order" edge.
func HasOrder() predicate.OrderValidation {
	return predicate.OrderValidation(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, OrderTable, OrderColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOrderWith applies the HasEdge predicate on the "order" edge with a given conditions (other predicates).
func HasOrderWith(preds ...predicate.Order) predicate.OrderValidation {
	return predicate.OrderValidation(func(s *sql.Selector) {
		step := newOrderStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.OrderValidation) predicate.OrderValidation {
	return predicate.OrderValidation(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.OrderValidation) predicate.OrderValidation {
	return predicate.OrderValidation(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.OrderValidation) predicate.OrderValidation {
	return predicate.OrderValidation(sql.NotPredicates(p))
}
