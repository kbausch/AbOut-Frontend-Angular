using AbOut_Database_Testing;
using Microsoft.VisualStudio.TestTools.UnitTesting;
using MySql.Data.MySqlClient;

namespace Outcomes_Testing
{
    [TestClass()]
    public class ViewOutcomes : TestConnectorMySQL
    {
        [TestMethod()]
        public void List_DefaultInput()
        {
            // Arrange:
            // Prepare the query for listing all outcomes.
            const string query = "SELECT prefix, identifier, text, begin, end " +
                "FROM outcomes__list_all__vw;";
            MySqlCommand cmd = new MySqlCommand();
            cmd.Connection = GetConnectionObject();
            cmd.CommandText = query;

            // Act:
            // Execute the query.
            using (MySqlDataReader result = cmd.ExecuteReader())
            {
                const int expectedFieldCount = 5;
                const int expectedRowCount = 3;
                string[,] desiredOutput = new string[expectedRowCount, expectedFieldCount]
                {
                    { "EAC","1","an ability to identify, formulate, and solve complex engineering problems by applying principles of engineering, science, and mathematics", "Fall 2016", "" },
                    { "EAC","2","an ability to apply engineering design to produce solutions that meet specified needs with consideration of public health, safety, and welfare, as well as global, cultural, social, environmental, and economic factors", "Fall 2016", "" },
                    { "CAC","1","an ability to analyze a complex computing problem and to apply principles of computing and other relevant disciplines to identify solutions", "Fall 2016", "" }
                };

                // Assert:
                // Compare expected with returned results.
                int rowCounter = 0;
                while (result.Read())
                {
                    // Ensure the proper amount of fields are returned.
                    Assert.AreEqual(expectedFieldCount, result.FieldCount);

                    // Ensure the row's fields contain the expected values.
                    Assert.AreEqual(desiredOutput[rowCounter, 0], result.GetString(0));
                    Assert.AreEqual(desiredOutput[rowCounter, 1], result.GetString(1));
                    Assert.AreEqual(desiredOutput[rowCounter, 2], result.GetString(2));
                    Assert.AreEqual(desiredOutput[rowCounter, 3], result.GetString(3));
                    Assert.AreEqual(desiredOutput[rowCounter, 4], result.GetString(4));

                    rowCounter++;
                }
            }
        }
    }
}