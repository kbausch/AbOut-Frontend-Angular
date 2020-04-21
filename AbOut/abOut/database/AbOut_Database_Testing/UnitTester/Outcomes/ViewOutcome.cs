using System;
using System.Diagnostics;
using AbOut_Database_Testing;
using Microsoft.VisualStudio.TestTools.UnitTesting;
using MySql.Data.MySqlClient;

namespace Outcomes_Testing
{
    [TestClass()]
    public class ViewOutcome : TestConnectorMySQL
    {
        [TestMethod()]
        public void Read_GoodInput()
        {
            // Arrange:
            // Prepare the query for fetching a single outcome.
            const string query = "CALL outcomes__list_one__sp('EAC', 1);";
            MySqlCommand cmd = new MySqlCommand();
            cmd.Connection = GetConnectionObject();
            cmd.CommandText = query;

            // Act:
            // Execute the query.
            using (MySqlDataReader result = cmd.ExecuteReader())
            {
                const int expectedFieldCount = 5;
                const int expectedRowCount = 1;
                string[,] desiredOutput = new string[expectedRowCount, expectedFieldCount]
                {
                    { "EAC","1","an ability to identify, formulate, and solve complex engineering problems by applying principles of engineering, science, and mathematics", "Fall 2016", "" },
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

                // Ensure the rows are of the expected length.
                Assert.AreEqual(expectedRowCount, rowCounter);

            }
        }

        [TestMethod()]
        public void Read_InvalidInput_PrefixNotFound()
        {
            // Arrange:
            // Prepare the query for fetching a single outcome.
            const string query = "CALL outcomes__list_one__sp('EA', 1);";
            MySqlCommand cmd = new MySqlCommand();
            cmd.Connection = GetConnectionObject();
            cmd.CommandText = query;

            // Act:
            // Execute the query.
            using (MySqlDataReader result = cmd.ExecuteReader())
            {
                // Assert:
                // Ensure the rows are empty as the query should be invalid.
                Assert.IsFalse(result.Read());
            }
        }

        [TestMethod()]
        public void Read_InvalidInput_PrefixNull()
        {
            // Arrange:
            // Prepare the query for fetching a single outcome.
            const string query = "CALL outcomes__list_one__sp(null, 1);";
            MySqlCommand cmd = new MySqlCommand();
            cmd.Connection = GetConnectionObject();
            cmd.CommandText = query;

            // Act:
            // Execute the query.
            using (MySqlDataReader result = cmd.ExecuteReader())
            {
                // Assert:
                // Ensure the rows are empty as the query should be invalid.
                Assert.IsFalse(result.Read());
            }
        }

        [TestMethod()]
        public void Read_InvalidInput_IdentifierNotFound()
        {
            // Arrange:
            // Prepare the query for fetching a single outcome.
            const string query = "CALL outcomes__list_one__sp('EAC', 4);";
            MySqlCommand cmd = new MySqlCommand();
            cmd.Connection = GetConnectionObject();
            cmd.CommandText = query;

            // Act:
            // Execute the query.
            using (MySqlDataReader result = cmd.ExecuteReader())
            {
                // Assert:
                // Ensure the rows are empty as the query should be invalid.
                Assert.IsFalse(result.Read());
            }
        }

        [TestMethod()]
        public void Read_InvalidInput_IdentifierNull()
        {
            // Arrange:
            // Prepare the query for fetching a single outcome.
            const string query = "CALL outcomes__list_one__sp('EAC', null);";
            MySqlCommand cmd = new MySqlCommand();
            cmd.Connection = GetConnectionObject();
            cmd.CommandText = query;

            // Act:
            // Execute the query.
            using (MySqlDataReader result = cmd.ExecuteReader())
            {
                // Assert:
                // Ensure the rows are empty as the query should be invalid.
                Assert.IsFalse(result.Read());
            }
        }
    }
}