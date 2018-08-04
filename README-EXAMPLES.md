The CLI will generate real world style data to submit via a schema. The schema will dictate basic structure from data types, how many elements of data to generate, and the particular kind of data. For example, data types would include types like integer, float, strings, and similar. The number of elements will simply be the number of items to generate in total. The kind of data would entail things like addresses, names, first names, last names, zip codes, credit cards, general lorum ipsum, and other kinds.

* GenerateCount - How many elements of data to generate.

* SourceDestination - Source destination of where the generated data would be sent. Some possible options (and hope to support) might include; file, console, SQL Server, Postgresql, ?? Maria DB ??, Cassandra, DataStax Enterprise, and [CosmosDB](https://azure.microsoft.com/en-us/services/cosmos-db/).

* Schema

{
    "GenerateCount": "10",
    "Destination": {
        "Source": "Cassandra",
        "Connection": "localhost:9042",
        "Auth": {
            "Type": "PlainText",
            "Username": "DontPutitHerePutItInVault",
            "Password": "SameGetItInVault"
        }
    },
    "Schema": {
        "Table": {
            "Name": "Passengers",
            "Columns": {
                "Passenger": "FullName",
                "FirstName": "FirstName",
                "LastName": "LastName",
                "DateOfBirth": "RandomDate",
                "Minor": "Bool",
                "Paid": "Money",
                "Cost": "Money",
                "Class": "Range"
            }
        }
    }
}

In this example, we would need direct generation for the elementes FullName, FirstName, and LastName. For DateOfBirth this value would need to be seeded as very low probabiliy for a minor's date of birth, but more likely for adaults of various demographics. Minor would be a calculated field off of Date of Birth. Then Paid and Cost would need to be a random value for calculations sake, but based on an initial seed value of fare price. Then class would be a set range of values that could or would be randomly selected, possibly also based on a weighted scale of sorts.
