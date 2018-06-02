> NOTE: Shifted this directly to the [README.md](README.md) file for now. Eventually will add features issues in the issues list and put together some real documentation. 

# Colligere CLI Application Feature Descriptions

## Description

The idea behind this CLI is to specifically create a way to send generated data against a particular API end point.

### Features

#### Data Generation

The CLI will generate real world style data to submit via a schema. The schema will dictate basic structure from data stypes, how many elements of data to generate, and the particular kind of data. For example, data types would include types like integer, float, strings, and similar. The number of elements will simply be the number of items to generate in total. The kind of data would entail things like addresses, names, first names, last names, zip codes, credit cards, general lorum ipsum, and other kinds.

#### End Point

The end point in which the CLI can submit to woul include a singular URI end point where the data payload would be sent. Using HTTP, with a verb - the default being POST - the data would be sent readily as a singular request.

### Future Feature Ideas & Thoughts

The submission of data could be changed to be generated and sent via an ongoing stream of data vs as a bulk clump of data for ingest.

Another feature may be to rate limit, or regulate the submission of data sent to the end point.

Threading/Go Routines could be added, and maybe even a parameter, to add how many go routines would be used to simultaneously send requests to the end point.
