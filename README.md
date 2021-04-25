# AboutGo

数据库层面:

    return errors.Warf(sql.ErrNoRows, "db: sql exec failed with no data")


业务层面:

    if errors.Is(err, sql.ErrNoRows) {
        // do something
    }
