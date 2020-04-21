using System;
using System.Data;
using AbOut_Database_Testing;
using Microsoft.VisualStudio.TestTools.UnitTesting;
using MySql.Data.MySqlClient;

namespace ProgramOutcomes_Testing
{

    [TestClass()]
    public class DisassociateOutcome : TestConnectorMySQL
    {
        [TestMethod()]
        public void Disassociate_GoodInput()
        {
            // Arrange:
            // Prepare the query for fetching a single outcome.
            MySqlCommand cmd = new MySqlCommand();
            cmd.Connection = GetConnectionObject();
            cmd.CommandType = System.Data.CommandType.StoredProcedure;
            cmd.CommandText = "program_outcomes__disassociate_outcome__sp";
            cmd.Parameters.Add("program_abbrev", MySqlDbType.VarChar).Value = "SE";
            cmd.Parameters.Add("prefix_text", MySqlDbType.VarChar).Value = "EAC";
            cmd.Parameters.Add("outcome_identifier", MySqlDbType.VarChar).Value = "1";
            cmd.Parameters.Add("status", MySqlDbType.Int32).Direction = ParameterDirection.Output;
            cmd.Parameters.Add("error_message", MySqlDbType.VarChar).Direction = ParameterDirection.Output;

            // We create and assign a transaction to the command.
            MySqlTransaction transaction = cmd.Connection.BeginTransaction();
            cmd.Transaction = transaction;

            try
            {
                // Act:
                // Execute the query.
                cmd.ExecuteNonQuery();

                // Store the result parameters.
                int status = Convert.ToInt32(cmd.Parameters["status"].Value);
                string errorMessage = Convert.ToString(cmd.Parameters["error_message"].Value);

                // Assert:
                // Ensure the status is 0 for no error.
                const int expectedStatus = 0;
                Assert.AreEqual(expectedStatus, status);

                // The error_message string should be empty if no problems occured.
                Assert.AreEqual("", errorMessage);

                // Ensure the association was removed.
                const string tquery = "SELECT id FROM program_outcomes " + 
                    "WHERE program_id = 1 AND outcome_id = 1 LIMIT 1;";
                cmd.CommandType = System.Data.CommandType.Text;
                cmd.CommandText = tquery;
                cmd.Parameters.Clear();
                string association = Convert.ToString(cmd.ExecuteScalar());
                Assert.AreEqual("", association);
            }
            finally
            {
                // Cleanup:
                // Remove the outcome we created.
                transaction.Rollback();
            }
        }

        [TestMethod()]
        public void Disassociate_InvalidInput_AssociationNotFound()
        {
            // Arrange:
            // Prepare the query for fetching a single outcome.
            MySqlCommand cmd = new MySqlCommand();
            cmd.Connection = GetConnectionObject();
            cmd.CommandType = System.Data.CommandType.StoredProcedure;
            cmd.CommandText = "program_outcomes__disassociate_outcome__sp";
            cmd.Parameters.Add("program_abbrev", MySqlDbType.VarChar).Value = "CS";
            cmd.Parameters.Add("prefix_text", MySqlDbType.VarChar).Value = "CAC";
            cmd.Parameters.Add("outcome_identifier", MySqlDbType.VarChar).Value = "1";
            cmd.Parameters.Add("status", MySqlDbType.Int32).Direction = ParameterDirection.Output;
            cmd.Parameters.Add("error_message", MySqlDbType.VarChar).Direction = ParameterDirection.Output;

            // We create and assign a transaction to the command.
            MySqlTransaction transaction = cmd.Connection.BeginTransaction();
            cmd.Transaction = transaction;

            try
            {
                // Act:
                // Execute the query.
                cmd.ExecuteNonQuery();

                // Store the result parameters.
                int status = Convert.ToInt32(cmd.Parameters["status"].Value);
                string errorMessage = Convert.ToString(cmd.Parameters["error_message"].Value);

                // Assert:
                // Ensure the status is 1 for an error.
                const int expectedStatus = 1;
                Assert.AreEqual(expectedStatus, status);

                Assert.AreEqual("Association does not exist.", errorMessage);
            }
            finally
            {
                // Cleanup:
                // Remove the outcome we created.
                transaction.Rollback();
            }
        }

        [TestMethod()]
        public void Disassociate_InvalidInput_ProgramAbbrevNotFound()
        {
            // Arrange:
            // Prepare the query for fetching a single outcome.
            MySqlCommand cmd = new MySqlCommand();
            cmd.Connection = GetConnectionObject();
            cmd.CommandType = System.Data.CommandType.StoredProcedure;
            cmd.CommandText = "program_outcomes__disassociate_outcome__sp";
            cmd.Parameters.Add("program_abbrev", MySqlDbType.VarChar).Value = "BRUH";
            cmd.Parameters.Add("prefix_text", MySqlDbType.VarChar).Value = "EAC";
            cmd.Parameters.Add("outcome_identifier", MySqlDbType.VarChar).Value = "1";
            cmd.Parameters.Add("status", MySqlDbType.Int32).Direction = ParameterDirection.Output;
            cmd.Parameters.Add("error_message", MySqlDbType.VarChar).Direction = ParameterDirection.Output;

            // We create and assign a transaction to the command.
            MySqlTransaction transaction = cmd.Connection.BeginTransaction();
            cmd.Transaction = transaction;

            try
            {
                // Act:
                // Execute the query.
                cmd.ExecuteNonQuery();

                // Store the result parameters.
                int status = Convert.ToInt32(cmd.Parameters["status"].Value);
                string errorMessage = Convert.ToString(cmd.Parameters["error_message"].Value);

                // Assert:
                // Ensure the status is 1 for an error.
                const int expectedStatus = 1;
                Assert.AreEqual(expectedStatus, status);

                Assert.AreEqual("Program does not exist.", errorMessage);
            }
            finally
            {
                // Cleanup:
                // Remove the outcome we created.
                transaction.Rollback();
            }
        }

        [TestMethod()]
        public void Disassociate_InvalidInput_PrefixNotFound()
        {
            // Arrange:
            // Prepare the query for fetching a single outcome.
            MySqlCommand cmd = new MySqlCommand();
            cmd.Connection = GetConnectionObject();
            cmd.CommandType = System.Data.CommandType.StoredProcedure;
            cmd.CommandText = "program_outcomes__disassociate_outcome__sp";
            cmd.Parameters.Add("program_abbrev", MySqlDbType.VarChar).Value = "SE";
            cmd.Parameters.Add("prefix_text", MySqlDbType.VarChar).Value = "DALTO";
            cmd.Parameters.Add("outcome_identifier", MySqlDbType.VarChar).Value = "1";
            cmd.Parameters.Add("status", MySqlDbType.Int32).Direction = ParameterDirection.Output;
            cmd.Parameters.Add("error_message", MySqlDbType.VarChar).Direction = ParameterDirection.Output;

            // We create and assign a transaction to the command.
            MySqlTransaction transaction = cmd.Connection.BeginTransaction();
            cmd.Transaction = transaction;

            try
            {
                // Act:
                // Execute the query.
                cmd.ExecuteNonQuery();

                // Store the result parameters.
                int status = Convert.ToInt32(cmd.Parameters["status"].Value);
                string errorMessage = Convert.ToString(cmd.Parameters["error_message"].Value);

                // Assert:
                // Ensure the status is 1 for an error.
                const int expectedStatus = 1;
                Assert.AreEqual(expectedStatus, status);

                Assert.AreEqual("Outcome does not exist.", errorMessage);
            }
            finally
            {
                // Cleanup:
                // Remove the outcome we created.
                transaction.Rollback();
            }
        }

        [TestMethod()]
        public void Disassociate_InvalidInput_OutcomeIdentifierNotFound()
        {
            // Arrange:
            // Prepare the query for fetching a single outcome.
            MySqlCommand cmd = new MySqlCommand();
            cmd.Connection = GetConnectionObject();
            cmd.CommandType = System.Data.CommandType.StoredProcedure;
            cmd.CommandText = "program_outcomes__disassociate_outcome__sp";
            cmd.Parameters.Add("program_abbrev", MySqlDbType.VarChar).Value = "SE";
            cmd.Parameters.Add("prefix_text", MySqlDbType.VarChar).Value = "EAC";
            cmd.Parameters.Add("outcome_identifier", MySqlDbType.VarChar).Value = "500";
            cmd.Parameters.Add("status", MySqlDbType.Int32).Direction = ParameterDirection.Output;
            cmd.Parameters.Add("error_message", MySqlDbType.VarChar).Direction = ParameterDirection.Output;

            // We create and assign a transaction to the command.
            MySqlTransaction transaction = cmd.Connection.BeginTransaction();
            cmd.Transaction = transaction;

            try
            {
                // Act:
                // Execute the query.
                cmd.ExecuteNonQuery();

                // Store the result parameters.
                int status = Convert.ToInt32(cmd.Parameters["status"].Value);
                string errorMessage = Convert.ToString(cmd.Parameters["error_message"].Value);

                // Assert:
                // Ensure the status is 1 for an error.
                const int expectedStatus = 1;
                Assert.AreEqual(expectedStatus, status);

                Assert.AreEqual("Outcome does not exist.", errorMessage);
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