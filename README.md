# Gohn

This repository is home to a Go client for the [Hacker News API][0].

## Name

The name of this project is a portmanteau of "Go" and "Hacker News" and is pronounced like "Gone" or "John."

## TODO

This file is where I keep track of what needs to be done for the library to be ready for a 1.0.0 release and be considered "feature complete."

- [ ] Export objects/structs for all API responses and objects.
- [ ] Provide methods for all top-level API endpoints.
- [ ] Provide helper methods to extend the functionality of the API.
- [ ] Add tests.
  - [ ] I've never written tests for Go before so this should be fun 😅

### Routes

| Route             | Method | Description                                      |
| ----------------- | ------ | ------------------------------------------------ |
| `/item/:itemId`   | `GET`  | Retrieve an individual item.                     |
| `/user/:username` | `GET`  | Retrieve an individual user.                     |
| `/maxitem`        | `GET`  | Get the current largest item ID.                 |
| `/topstories`     | `GET`  | Get up to 500 of the top stories.                |
| `/newstories`     | `GET`  | Get up to 500 of the newest stories.             |
| `/beststories`    | `GET`  | Get up to 500 of the best stories.               |
| `/askstories`     | `GET`  | Get up to 200 of the latest **Ask HN** stories.  |
| `/showstories`    | `GET`  | Get up to 200 of the latest **Show HN** stories. |
| `/jobstories`     | `GET`  | Get up to 200 of the latest **Job** stories.     |
| `/updates`        | `GET`  | List of recent item and profile changes.         |



[0]: https://github.com/HackerNews/API
