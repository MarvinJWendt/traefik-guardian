package api

// func Authenticate(store *session.Store) func(ctx *fiber.Ctx) error {
// 	return func(ctx *fiber.Ctx) error {
// 		password := ctx.Get("password")
//
// 		if !auth.CheckPassword(password) {
// 			return ctx.SendStatus(fiber.StatusUnauthorized)
// 		}
//
// 		sess, err := store.Get(ctx)
// 		if err != nil {
// 			return err
// 		}
//
// 		err = auth.Authenticate(sess)
// 		if err != nil {
// 			return err
// 		}
//
// 		return ctx.SendStatus(fiber.StatusOK)
// 	}
// }
