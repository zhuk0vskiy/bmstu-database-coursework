
func TestStudioPostrgresql_Get(t *testing.T) {
	type fields struct {
		db *pgxpool.Pool
	}
	type args struct {
		ctx     context.Context
		request *dto.GetStudioRequest
	}
	tests := []struct {
		name string
		//fields     fields
		args       args
		wantStudio *model.Studio
		wantErr    bool
	}{
		// TODO: Add test cases.
		{
			name: "test_pos_01",
			args: args{

				ctx: context.Background(),
				request: &dto.GetStudioRequest{
					Id: 1,
				},
			},
			wantStudio: &model.Studio{
				Id:   1,
				Name: "first",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := postgresql.NewStudioRepository(testDbInstance)

			gotStudio, err := p.Get(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotStudio, tt.wantStudio) {
				t.Errorf("Get() gotStudio = %v, want %v", gotStudio, tt.wantStudio)
			}
		})
	}
}
