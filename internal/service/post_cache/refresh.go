package post_cache

func (s *service) Refresh(limit int64) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	var err error
	s.requests, err = s.postroutingRepo.GetRequests(limit)
	if err != nil {
		return err
	}
	s.offset = 0

	return nil
}
