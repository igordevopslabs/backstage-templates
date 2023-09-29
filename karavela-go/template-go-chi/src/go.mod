module ${{ values.apiName | lower }}

go ${{ values.goVersion | lower }}

require github.com/go-chi/chi v1.5.5
