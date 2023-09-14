package room

func ConvertToRoomResponse(r Room) RoomResponse{
	return RoomResponse{
		ID : r.ID,
		Name : r.Name,
		Capacity : r.Capacity,
		Status : r.Status,
		CreatedAt : r.CreatedAt,
		UpdatedAt : r.UpdatedAt,
	}
}