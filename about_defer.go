package koans

func aboutDefer() {
	func() {
		i := 0
		defer func() {
			i++
			assert(i == intintint)
		}()
		i++
		assert(i == intintint)
	}()

	// multiple defered. LIFO? FIFO? or other?
	func() {
		i := 0
		defer func() {
			i++
			assert(i == intintint)
		}()
		defer func() {
			i += 2
			assert(i == intintint)
		}()
		defer func() {
			i += 3
			assert(i == intintint)
		}()

		assert(i == intintint)
	}()
}
