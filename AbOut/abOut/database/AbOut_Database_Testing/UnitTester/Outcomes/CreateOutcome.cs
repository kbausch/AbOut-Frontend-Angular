using System;
using System.Data;
using AbOut_Database_Testing;
using Microsoft.VisualStudio.TestTools.UnitTesting;
using MySql.Data.MySqlClient;

namespace Outcomes_Testing
{
    [TestClass()]
    public class CreateOutcome : TestConnectorMySQL
    {
        [TestMethod()]
        public void Create_GoodInput()
        {
            // Arrange:
            // Prepare the query for fetching a single outcome.
            MySqlCommand cmd = new MySqlCommand();
            cmd.Connection = GetConnectionObject();
            cmd.CommandType = System.Data.CommandType.StoredProcedure;
            cmd.CommandText = "outcomes__create_outcome__sp";
            cmd.Parameters.Add("pre", MySqlDbType.VarChar).Value = "EAC";
            cmd.Parameters.Add("idnt", MySqlDbType.VarChar).Value = "5";
            cmd.Parameters.Add("txt", MySqlDbType.VarChar).Value = "One test outcome created.";
            cmd.Parameters.Add("status", MySqlDbType.Int32).Direction = ParameterDirection.Output;
            cmd.Parameters.Add("error_message", MySqlDbType.VarChar).Direction = ParameterDirection.Output;

            // We create and assign a transaction to the command.
            MySqlTransaction transaction = cmd.Connection.BeginTransaction();
            cmd.Transaction = transaction;

            // Act:
            // Execute the query.
            MySqlDataReader result = cmd.ExecuteReader();

            try
            {
                // Store the result parameters.
                int status = Convert.ToInt32(cmd.Parameters["status"].Value);
                string errorMessage = Convert.ToString(cmd.Parameters["error_message"].Value);

                // Assert:
                // Ensure the status is 0 for no error.
                const int expectedStatus = 0;
                Assert.AreEqual(expectedStatus, status);

                // The error_message string should be empty if no problems occured.
                Assert.AreEqual("", errorMessage);
            }
            finally
            {
                result.Close();

                // Cleanup:
                // Remove the outcome association we created.
                transaction.Rollback();
            }
        }

        [TestMethod()]
        public void Create_InvalidInput_PrefixNotFound()
        {
            // Arrange:
            // Prepare the query for fetching a single outcome.
            MySqlCommand cmd = new MySqlCommand();
            cmd.Connection = GetConnectionObject();
            cmd.CommandType = System.Data.CommandType.StoredProcedure;
            cmd.CommandText = "outcomes__create_outcome__sp";
            cmd.Parameters.Add("pre", MySqlDbType.VarChar).Value = "BRUH";
            cmd.Parameters.Add("idnt", MySqlDbType.VarChar).Value = "5";
            cmd.Parameters.Add("txt", MySqlDbType.VarChar).Value = "One test outcome created.";
            cmd.Parameters.Add("status", MySqlDbType.Int32).Direction = ParameterDirection.Output;
            cmd.Parameters.Add("error_message", MySqlDbType.VarChar).Direction = ParameterDirection.Output;

            // Act:
            // Execute the query.
            MySqlDataReader result = cmd.ExecuteReader();

            try
            {
                // Store the result parameters.
                int status = Convert.ToInt32(cmd.Parameters["status"].Value);
                string errorMessage = Convert.ToString(cmd.Parameters["error_message"].Value);

                // Assert:
                // Ensure the status is 1 for error.
                const int expectedStatus = 1;
                Assert.AreEqual(expectedStatus, status);

                // The error_message should alert us of the problem.
                Assert.AreEqual("Prefix does not exist", errorMessage);
            }
            finally
            {
                result.Close();
            }
        }

        [TestMethod()]
        public void Create_InvalidInput_PrefixIdentifierPairAlreadyExists()
        {
            // Arrange:
            // Prepare the query for fetching a single outcome.
            MySqlCommand cmd = new MySqlCommand();
            cmd.Connection = GetConnectionObject();
            cmd.CommandType = System.Data.CommandType.StoredProcedure;
            cmd.CommandText = "outcomes__create_outcome__sp";
            cmd.Parameters.Add("pre", MySqlDbType.VarChar).Value = "EAC";
            cmd.Parameters.Add("idnt", MySqlDbType.VarChar).Value = "1";
            cmd.Parameters.Add("txt", MySqlDbType.VarChar).Value = "One test outcome created.";
            cmd.Parameters.Add("status", MySqlDbType.Int32).Direction = ParameterDirection.Output;
            cmd.Parameters.Add("error_message", MySqlDbType.VarChar).Direction = ParameterDirection.Output;

            // Act:
            // Execute the query.
            MySqlDataReader result = cmd.ExecuteReader();

            try
            {
                // Store the result parameters.
                int status = Convert.ToInt32(cmd.Parameters["status"].Value);
                string errorMessage = Convert.ToString(cmd.Parameters["error_message"].Value);

                // Assert:
                // Ensure the status is 1 for error.
                const int expectedStatus = 1;
                Assert.AreEqual(expectedStatus, status);

                // The error_message should alert us of the problem.
                Assert.AreEqual("Outcome already exists with this prefix and identifier", errorMessage);
            }
            finally
            {
                result.Close();
            }
        }

    }
}