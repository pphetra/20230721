package member_infra

import (
	"database/sql"
	member_domain "taejai/internal/member/domain"
	shared_domain "taejai/internal/shared/domain"
)

type MemberRepository struct {
	tx *sql.Tx
}

func NewMemberRepository(tx *sql.Tx) *MemberRepository {
	return &MemberRepository{tx: tx}
}

func (r *MemberRepository) GetById(id member_domain.MemberId) (*member_domain.Member, error) {
	member := &member_domain.Member{}
	err := r.tx.QueryRow("SELECT id, name1, name2, member_type, email, address_line1, address_line2, address_postal_code FROM members WHERE id = $1", id).Scan(
		&member.Id,
		&member.Name1,
		&member.Name2,
		&member.Type,
		&member.Email,
		&member.Address.Line1,
		&member.Address.Line2,
		&member.Address.PostalCode,
	)
	if err != nil {
		return nil, err
	}
	return member, nil
}

func (r *MemberRepository) FindByName(name string) ([]*member_domain.Member, error) {
	rows, err := r.tx.Query("SELECT id, name1, name2, member_type, email, address_line1, address_line2, address_postal_code FROM members WHERE name1 LIKE $1 OR name2 LIKE $1", "%"+name+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	members := []*member_domain.Member{}
	for rows.Next() {
		member := &member_domain.Member{}
		err := rows.Scan(
			&member.Id,
			&member.Name1,
			&member.Name2,
			&member.Type,
			&member.Email,
			&member.Address.Line1,
			&member.Address.Line2,
			&member.Address.PostalCode,
		)
		if err != nil {
			return nil, err
		}
		members = append(members, member)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return members, nil
}

func (r *MemberRepository) Create(member *member_domain.Member) (member_domain.MemberId, error) {
	var id int
	err := r.tx.QueryRow(
		"INSERT INTO members (name1, name2, member_type, email, address_line1, address_line2, address_postal_code) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id",
		member.Name1,
		member.Name2,
		member.Type,
		member.Email,
		member.Address.Line1,
		member.Address.Line2,
		member.Address.PostalCode,
	).Scan(&id)
	if err != nil {
		return 0, err
	}
	return member_domain.MemberId(id), nil
}

func (r *MemberRepository) Update(member *member_domain.Member) error {
	_, err := r.tx.Exec(
		"UPDATE members SET name1 = $1, name2 = $2, member_type = $3, email = $4, address_line1 = $5, address_line2 = $6, address_postal_code = $7 WHERE id = $8",
		member.Name1,
		member.Name2,
		member.Type,
		member.Email,
		member.Address.Line1,
		member.Address.Line2,
		member.Address.PostalCode,
		member.Id,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *MemberRepository) Delete(member *member_domain.Member) error {
	_, err := r.tx.Exec("DELETE FROM members WHERE id = $1", member.Id)
	if err != nil {
		return err
	}
	return nil
}

// =======================
func init() {
	shared_domain.RepositoryRegistry.Register("member", func(tx *sql.Tx) interface{} {
		return NewMemberRepository(tx)
	})
}
