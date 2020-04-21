using System;
using System.Data;
using AbOut_Database_Testing;
using Microsoft.VisualStudio.TestTools.UnitTesting;
using MySql.Data.MySqlClient;

namespace ProgramOutcomes_Testing
{

    [TestClass()]
    public class AssociateOutcome : TestConnectorMySQL
    {
        [TestMethod()]
        public void Associate_GoodInput()
        {
            // Arrange:
            // Prepare the query for fetching a single outcome.
            MySqlCommand cmd = new MySqlCommand();
            cmd.Connection = GetConnectionObject();
            cmd.CommandType = System.Data.CommandType.StoredProcedure;
            cmd.CommandText = "program_outcomes__associate_outcome__sp";
            cmd.Parameters.Add("program_abbrev", MySqlDbType.VarChar).Value = "CS";
            cmd.Parameters.Add("prefix_text", MySqlDbType.VarChar).Value = "EAC";
            cmd.Parameters.Add("outcome_identifier", MySqlDbType.VarChar).Value = "1";
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
                // Remove the outcome we created.
                transaction.Rollback();
            }
        }

        [TestMethod()]
        public void Associate_GoodInput_NoInterval()
        {
            // Arrange:
            // Prepare the query for fetching a single outcome.
            MySqlCommand cmd = new MySqlCommand();
            cmd.Connection = GetConnectionObject();
            cmd.CommandType = System.Data.CommandType.StoredProcedure;
            cmd.CommandText = "program_outcomes__associate_outcome__sp";
            cmd.Parameters.Add("program_abbrev", MySqlDbType.VarChar).Value = "CS";
            cmd.Parameters.Add("prefix_text", MySqlDbType.VarChar).Value = "CAC";
            cmd.Parameters.Add("outcome_identifier", MySqlDbType.VarChar).Value = "1";
            cmd.Parameters.Add("status", MySqlDbType.Int32).Direction = ParameterDirection.Output;
            cmd.Parameters.Add("error_message", MySqlDbType.VarChar).Direction = ParameterDirection.Output;

            // We create and assign a transaction to the command.
            MySqlTransaction transaction = cmd.Connection.BeginTransaction();
            cmd.Transaction = transaction;

            // Act:
            // Execute the query.
            cmd.ExecuteNonQuery();

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
                // Cleanup:
                // Remove the outcome we created.
                transaction.Rollback();
            }
        }

        [TestMethod()]
        public void Associate_InvalidInput_AssociationAlreadyExists()
        {
            // Arrange:
            // Prepare the query for fetching a single outcome.
            MySqlCommand cmd = new MySqlCommand();
            cmd.Connection = GetConnectionObject();
            cmd.CommandType = System.Data.CommandType.StoredProcedure;
            cmd.CommandText = "program_outcomes__associate_outcome__sp";
            cmd.Parameters.Add("program_abbrev", MySqlDbType.VarChar).Value = "SE";
            cmd.Parameters.Add("prefix_text", MySqlDbType.VarChar).Value = "EAC";
            cmd.Parameters.Add("outcome_identifier", MySqlDbType.VarChar).Value = "1";
            cmd.Parameters.Add("status", MySqlDbType.Int32).Direction = ParameterDirection.Output;
            cmd.Parameters.Add("error_message", MySqlDbType.VarChar).Direction = ParameterDirection.Output;

            // We create and assign a transaction to the command.
            MySqlTransaction transaction = cmd.Connection.BeginTransaction();
            cmd.Transaction = transaction;

            // Act:
            // Execute the query.
            cmd.ExecuteNonQuery();

            try
            {
                // Store the result parameters.
                int status = Convert.ToInt32(cmd.Parameters["status"].Value);
                string errorMessage = Convert.ToString(cmd.Parameters["error_message"].Value);

                // Assert:
                // Ensure the status is 1 for no error.
                const int expectedStatus = 1;
                Assert.AreEqual(expectedStatus, status);

                // The error_message string should be empty if no problems occured.
                Assert.AreEqual("association already exists", errorMessage);
            }
            finally
            {
                // Cleanup:
                // Remove the outcome we created.
                transaction.Rollback();
            }
        }

        [TestMethod()]
        public void Associate_InvalidInput_OutcomeNotFound()
        {
            // Arrange:
            // Prepare the query for fetching a single outcome.
            MySqlCommand cmd = new MySqlCommand();
            cmd.Connection = GetConnectionObject();
            cmd.CommandType = System.Data.CommandType.StoredProcedure;
            cmd.CommandText = "program_outcomes__associate_outcome__sp";
            cmd.Parameters.Add("program_abbrev", MySqlDbType.VarChar).Value = "SE";
            cmd.Parameters.Add("prefix_text", MySqlDbType.VarChar).Value = "EAC";
            cmd.Parameters.Add("outcome_identifier", MySqlDbType.VarChar).Value = "420";
            cmd.Parameters.Add("status", MySqlDbType.Int32).Direction = ParameterDirection.Output;
            cmd.Parameters.Add("error_message", MySqlDbType.VarChar).Direction = ParameterDirection.Output;

            // We create and assign a transaction to the command.
            MySqlTransaction transaction = cmd.Connection.BeginTransaction();
            cmd.Transaction = transaction;

            // Act:
            // Execute the query.
            cmd.ExecuteNonQuery();

            try
            {
                // Store the result parameters.
                int status = Convert.ToInt32(cmd.Parameters["status"].Value);
                string errorMessage = Convert.ToString(cmd.Parameters["error_message"].Value);

                // Assert:
                // Ensure the status is 1 for no error.
                const int expectedStatus = 1;
                Assert.AreEqual(expectedStatus, status);

                // The error_message string should be empty if no problems occured.
                Assert.AreEqual("outcome does not exist", errorMessage);
            }
            finally
            {
                // Cleanup:
                // Remove the outcome we created.
                transaction.Rollback();
            }
        }

        [TestMethod()]
        public void Associate_InvalidInput_ProgramNotFound()
        {
            // Arrange:
            // Prepare the query for fetching a single outcome.
            MySqlCommand cmd = new MySqlCommand();
            cmd.Connection = GetConnectionObject();
            cmd.CommandType = System.Data.CommandType.StoredProcedure;
            cmd.CommandText = "program_outcomes__associate_outcome__sp";
            cmd.Parameters.Add("program_abbrev", MySqlDbType.VarChar).Value = "SUGMA";
            cmd.Parameters.Add("prefix_text", MySqlDbType.VarChar).Value = "EAC";
            cmd.Parameters.Add("outcome_identifier", MySqlDbType.VarChar).Value = "1";
            cmd.Parameters.Add("status", MySqlDbType.Int32).Direction = ParameterDirection.Output;
            cmd.Parameters.Add("error_message", MySqlDbType.VarChar).Direction = ParameterDirection.Output;

            // We create and assign a transaction to the command.
            MySqlTransaction transaction = cmd.Connection.BeginTransaction();
            cmd.Transaction = transaction;

            // Act:
            // Execute the query.
            cmd.ExecuteNonQuery();

            try
            {
                // Store the result parameters.
                int status = Convert.ToInt32(cmd.Parameters["status"].Value);
                string errorMessage = Convert.ToString(cmd.Parameters["error_message"].Value);

                // Assert:
                // Ensure the status is 1 for no error.
                const int expectedStatus = 1;
                Assert.AreEqual(expectedStatus, status);

                // The error_message string should be empty if no problems occured.
                Assert.AreEqual("program does not exist", errorMessage);
            }
            finally
            {
                // Cleanup:
                // Remove the outcome we created.
                transaction.Rollback();
            }
        }

    }

}