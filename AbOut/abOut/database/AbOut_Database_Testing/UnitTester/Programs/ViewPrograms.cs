using AbOut_Database_Testing;
using Microsoft.VisualStudio.TestTools.UnitTesting;
using MySql.Data.MySqlClient;

namespace Programs_Testing
{
    [TestClass()]
    public class ViewPrograms : TestConnectorMySQL
    {
        [TestMethod()]
        public void ListCurrent_DefaultInput()
        {
            // Arrange:
            // Prepare the query for listing all outcomes.
            const string query = "SELECT abbrev, name, `current semester` " +
                "FROM programs__list_current__vw;";
            MySqlCommand cmd = new MySqlCommand();
            cmd.Connection = GetConnectionObject();
            cmd.CommandText = query;

            // Act:
            // Execute the query.
            using (MySqlDataReader result = cmd.ExecuteReader())
            {
                const int expectedFieldCount = 3;
                const int expectedRowCount = 3;
                string[,] desiredOutput = new string[expectedRowCount, expectedFieldCount]
                {
                    { "SE","Software Engineering", "Spring 2020" },
                    { "CS","Computer Science", "Spring 2020" },
                    { "EE","Electrical Engineering", "Spring 2020" }
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

                    rowCounter++;
                }

                // Ensure the rows are of the expected length.
                Assert.AreEqual(expectedRowCount, rowCounter);
            }
        }
        
    }
}