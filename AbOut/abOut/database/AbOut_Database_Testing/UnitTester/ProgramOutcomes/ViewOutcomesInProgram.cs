using System;
using System.Data;
using AbOut_Database_Testing;
using Microsoft.VisualStudio.TestTools.UnitTesting;
using MySql.Data.MySqlClient;

namespace ProgramOutcomes_Testing
{

    [TestClass()]
    public class ViewOutcomesInPrograms : TestConnectorMySQL
    {
        [TestMethod()]
        public void ViewOutcomesInProgram_GoodInput()
        {
            // Arrange:
            // Prepare the query for fetching a single outcome.
            MySqlCommand cmd = new MySqlCommand();
            cmd.Connection = GetConnectionObject();
            cmd.CommandType = System.Data.CommandType.StoredProcedure;
            cmd.CommandText = "programs_outcomes__outcomes_in_program__sp";
            cmd.Parameters.Add("program_abbrev", MySqlDbType.VarChar).Value = "CS";
            cmd.Parameters.Add("status", MySqlDbType.Int32).Direction = ParameterDirection.Output;
            cmd.Parameters.Add("error_message", MySqlDbType.VarChar).Direction = ParameterDirection.Output;
            cmd.Parameters.AddWithValue("result", new DataTable()).Direction = ParameterDirection.ReturnValue;

            // Act:
            // Execute the query.
            cmd.ExecuteNonQuery();

            const int expectedRows = 1;
            const int expectedFields = 3;
            string[,] desiredOutput = new string[expectedRows, expectedFields]
            {
                { "EAC", "1", "an ability to identify, formulate, and solve complex engineering problems by applying principles of engineering, science, and mathematics" }
            };

            // Assert:
            // Ensure the proper row output was returned.
            DataTable data = (DataTable) cmd.Parameters["result"].Value;

            foreach (DataRow row in data.AsEnumerable())
            {
                Assert.AreEqual(desiredOutput[0, 0], row["prefix"]);
                Assert.AreEqual(desiredOutput[0, 1], row["identifier"]);
                Assert.AreEqual(desiredOutput[0, 2], row["text"]);
            }

            // Store the result parameters.
            int status = Convert.ToInt32(cmd.Parameters["status"].Value);
            string errorMessage = Convert.ToString(cmd.Parameters["error_message"].Value);

            // Ensure the status is 0 for no error.
            const int expectedStatus = 0;
            Assert.AreEqual(expectedStatus, status);

            // The error_message string should be empty if no problems occured.
            Assert.AreEqual("", errorMessage);
        }

        [TestMethod()]
        public void ViewOutcomesInProgram_ProgramAbbrevNotFound()
        {
            // Arrange:
            // Prepare the query for fetching a single outcome.
            MySqlCommand cmd = new MySqlCommand();
            cmd.Connection = GetConnectionObject();
            cmd.CommandType = System.Data.CommandType.StoredProcedure;
            cmd.CommandText = "programs_outcomes__outcomes_in_program__sp";
            cmd.Parameters.Add("program_abbrev", MySqlDbType.VarChar).Value = "CAS";
            cmd.Parameters.Add("status", MySqlDbType.Int32).Direction = ParameterDirection.Output;
            cmd.Parameters.Add("error_message", MySqlDbType.VarChar).Direction = ParameterDirection.Output;

            // Act:
            // Execute the query.
            cmd.ExecuteNonQuery();

            // Assert:
            int status = Convert.ToInt32(cmd.Parameters["status"].Value);
            string errorMessage = Convert.ToString(cmd.Parameters["error_message"].Value);

            // Ensure the status is 1 for an error.
            const int expectedStatus = 1;
            Assert.AreEqual(expectedStatus, status);

            // The error_message string should be empty if no problems occured.
            Assert.AreEqual("Given program is invalid", errorMessage);
        }

    }

}