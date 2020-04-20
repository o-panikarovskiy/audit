package defusersrv

func (s *userService) destroyAuthSession(userID string) error {
	sid, err := s.sessions.Get(userID)

	if err != nil {
		return err
	}

	if sid != "" {
		_, err = s.sessions.Delete(sid)
	}

	if err != nil {
		return err
	}

	return nil
}