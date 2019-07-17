package rsschannel

import "database/sql"

func Fetch(channels *sql.Rows, ret *[]*Schema) (err error) {
	for channels.Next() {
		var channel Schema
		var title, description, link sql.NullString
		if err = channels.Scan(&channel.Id, &title, &description, &link, &channel.RSSLink); err != nil {
			return err
		}

		if title.Valid {
			channel.Title = title.String
		} else {
			channel.Title = " "
		}

		if description.Valid {
			channel.Description = description.String
		} else {
			channel.Description = " "
		}

		if link.Valid {
			channel.Link = link.String
		} else {
			channel.Link = " "
		}

		*ret = append(*ret, &channel)
	}
	return nil
}