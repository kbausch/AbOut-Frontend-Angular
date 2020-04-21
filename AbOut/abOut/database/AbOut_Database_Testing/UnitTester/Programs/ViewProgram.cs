using System;
using System.Data;
using AbOut_Database_Testing;
using Microsoft.VisualStudio.TestTools.UnitTesting;
using MySql.Data.MySqlClient;

namespace Programs_Testing
{
    [TestClass()]
    public class ViewProgram : TestConnectorMySQL
    {
        [TestMethod()]
        public void Read_GoodInput()
        {
            // Arrange:
            // Prepare the query for fetching a single program.
            MySqlCommand cmd = new MySqlCommand();
            cmd.Connection = GetConnectionObject();
            cmd.CommandType = System.Data.CommandType.StoredProcedure;
            cmd.CommandText = "programs__list_one__sp";
            cmd.Parameters.Add("abbrev", MySqlDbType.VarChar).Value = "CS";
            cmd.Parameters.Add("status", MySqlDbType.Int32).Direction = ParameterDirection.Output;
            cmd.Parameters.Add("error_message", MySqlDbType.VarChar).Direction = ParameterDirection.Output;
            cmd.Parameters.AddWithValue("result", new DataTable()).Direction = ParameterDirection.ReturnValue;

            // Act:
            // Execute the query.
            cmd.ExecuteNonQuery();

            const int expectedFieldCount = 3;
            const int expectedRowCount = 1;
            string[,] desiredOutput = new string[expectedRowCount, expectedFieldCount]
            {
                { "CS","Computer Science", "Spring 2020" },
            };

            // Assert:
            // Compare expected with returned results.

            // Ensure the row's fields contain the expected values.
            DataTable data = (DataTable) cmd.Parameters["result"].Value;

            foreach (DataRow row in data.AsEnumerable())
            {
                Assert.AreEqual(desiredOutput[0,0], row["abbrev"]);
                Assert.AreEqual(desiredOutput[0,1], row["name"]);
                Assert.AreEqual(desiredOutput[0,2], row["current_semester"]);
            }

            int status = Convert.ToInt32(cmd.Parameters["status"].Value);
            string errorMessage = Convert.ToString(cmd.Parameters["error_message"].Value);

            const int expectedStatus = 0;
            Assert.AreEqual(expectedStatus, status);

            Assert.AreEqual("", errorMessage);
        }

        [TestMethod()]
        public void Read_InvalidInput_AbbrevNotFound()
        {
            // Arrange:
            // Prepare the query for fetching a single program.
            MySqlCommand cmd = new MySqlCommand();
            cmd.Connection = GetConnectionObject();
            cmd.CommandType = System.Data.CommandType.StoredProcedure;
            cmd.CommandText = "programs__list_one__sp";
            cmd.Parameters.Add("abbrev", MySqlDbType.VarChar).Value = "CSSa";
            cmd.Parameters.Add("status", MySqlDbType.Int32).Direction = ParameterDirection.Output;
            cmd.Parameters.Add("error_message", MySqlDbType.VarChar).Direction = ParameterDirection.Output;
            cmd.Parameters.AddWithValue("result", new DataTable()).Direction = ParameterDirection.ReturnValue;

            // Act:
            // Execute the query.
            cmd.ExecuteNonQuery();

            // Assert:
            // Compare expected with returned results.

            int status = Convert.ToInt32(cmd.Parameters["status"].Value);
            string errorMessage = Convert.ToString(cmd.Parameters["error_message"].Value);

            const int expectedStatus = 1;
            Assert.AreEqual(expectedStatus, status);

            Assert.AreEqual("unknown program abbreviation", errorMessage);
        }

        [TestMethod()]
        public void Read_InvalidInput_AbbrevNull()
        {
            // Arrange:
            // Prepare the query for fetching a single program.
            MySqlCommand cmd = new MySqlCommand();
            cmd.Connection = GetConnectionObject();
            cmd.CommandType = System.Data.CommandType.StoredProcedure;
            cmd.CommandText = "programs__list_one__sp";
            cmd.Parameters.Add("abbrev", MySqlDbType.VarChar).Value = null;
            cmd.Parameters.Add("status", MySqlDbType.Int32).Direction = ParameterDirection.Output;
            cmd.Parameters.Add("error_message", MySqlDbType.VarChar).Direction = ParameterDirection.Output;
            cmd.Parameters.AddWithValue("result", new DataTable()).Direction = ParameterDirection.ReturnValue;

            // Act:
            // Execute the query.
            cmd.ExecuteNonQuery();

            // Assert:
            // Compare expected with returned results.

            int status = Convert.ToInt32(cmd.Parameters["status"].Value);
            string errorMessage = Convert.ToString(cmd.Parameters["error_message"].Value);

            const int expectedStatus = 1;
            Assert.AreEqual(expectedStatus, status);

            Assert.AreEqual("unknown program abbreviation", errorMessage);
        }

    }
}