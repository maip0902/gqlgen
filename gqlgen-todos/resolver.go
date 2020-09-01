package main

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	ssss []*Sixsixsix
}

// ↓ここ実装した
func (r *mutationResolver) Create666(ctx context.Context, input NewSixsixsix) (*Sixsixsix, error) {
    s := &Sixsixsix{
        Text:   input.Text,
        ID:     fmt.Sprintf("T%d", rand.Int()),
    }
    r.ssss = append(r.ssss, s)
    return s, nil
}

// ↓ここ実装した
func (r *queryResolver) Ssss(ctx context.Context) ([]*Sixsixsix, error) {
    return r.ssss, nil
}