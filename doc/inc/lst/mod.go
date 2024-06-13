
func TestEquipmentService_Get(t *testing.T) {
	type fields struct {
		equipmentRepo _interface.IEquipmentRepository
		reserveRepo   _interface.IReserveRepository
	}
	type args struct {
		request *dto.GetEquipmentRequest
	}
	tests := []struct {
		name          string
		fields        fields
		args          args
		wantEquipment *model.Equipment
		wantErr       bool
	}{
		{
			name: "test_pos_01",
			args: args{
				&dto.GetEquipmentRequest{
					Id: 1,
				},
			},
			wantErr: false,
			wantEquipment: &model.Equipment{
				Id:            1,
				Name:          "1",
				StudioId:      1,
				EquipmentType: 1,
			},
		},
		{
			name: "test_neg_01",
			args: args{
				&dto.GetEquipmentRequest{
					Id: 0,
				},
			},
			wantErr:       true,
			wantEquipment: nil,
		},
	}
	for _, tt := range tests {
		prodRepo := new(mocks.IEquipmentRepository)
		prodRepo.On("Get", context.Background(), tt.args.request).Return(&model.Equipment{
			Id:            1,
			Name:          "1",
			StudioId:      1,
			EquipmentType: 1,
		}, nil)
		t.Run(tt.name, func(t *testing.T) {
			s := EquipmentService{
				equipmentRepo: prodRepo,
				//reserveRepo:  tt.fields.reserveRepo,
			}
			gotEquipment, err := s.Get(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotEquipment, tt.wantEquipment) {
				t.Errorf("Get() gotEquipment = %v, want %v", gotEquipment, tt.wantEquipment)
			}
		})
	}
}
