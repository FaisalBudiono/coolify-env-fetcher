# Important!
Make sure the ENV you want to fetch does not have any sensitive data. This action only fetched the ENV that has a `build flag` on it, so make sure your `build flag` ENV is not sensitive.

# Usage Example

```yaml
name: Deploy

on:
  push:

jobs:
  build:
    name: Build
    runs-on: ubuntu-latest
    steps:
      - name: Fetch ENV from Coolify
        id: fetcher
        uses: FaisalBudiono/coolify-env-fetcher@v1
        with:
          base-url: ${{ secrets.COOLIFY_URL }}
          access-token: ${{ secrets.COOLIFY_TOKEN }}
          app-id: ${{ secrets.APP_ID }}
      - name: Upload .env
        uses: actions/upload-artifact@v4
        with:
          name: dev-env
          path: .env
          retention-days: 1
          include-hidden-files: true

      - name: Checkout
        uses: actions/checkout@v4
      - name: Download .env
        uses: actions/download-artifact@v4
        with:
          name: dev-env 
      - run: ls -la
```
