using System;
using System.Data;
using System.Diagnostics;
using AbOut_Database_Testing;
using Microsoft.VisualStudio.TestTools.UnitTesting;
using MySql.Data.MySqlClient;

namespace Outcomes_Testing
{
    [TestClass()]
    public class DeleteOutcome : TestConnectorMySQL
    {
        [TestMethod()]
        public void Delete_ValidDelete()
        {
            // Assert:
            // Prepare the delete command and transaction.

            // First, we create the delete command.
            MySqlCommand cmd = new MySqlCommand();
            cmd.Connection = GetConnectionObject();
            cmd.CommandType = System.Data.CommandType.StoredProcedure;
            cmd.CommandText = "outcomes__delete_outcome__sp";
            cmd.Parameters.Add("pre", MySqlDbType.VarChar).Value = "CAC";
            cmd.Parameters.Add("idnt", MySqlDbType.VarChar).Value = "1";
            cmd.Parameters.Add("status", MySqlDbType.Int32).Direction = ParameterDirection.Output;
            cmd.Parameters.Add("error_message", MySqlDbType.VarChar).Direction = ParameterDirection.Output;

            // Then, we create and assign a transaction to the command.
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
                // Undo the delete by rolling back the transaction.
                transaction.Rollback();
            }
        }

        [TestMethod()]
        public void Delete_Invalid_HasAssociation()
        {
            // Assert:
            // Prepare the delete command and transaction.

            // First, we create the delete command.
            MySqlCommand cmd = new MySqlCommand();
            cmd.Connection = GetConnectionObject();
            cmd.CommandType = System.Data.CommandType.StoredProcedure;
            cmd.CommandText = "outcomes__delete_outcome__sp";
            cmd.Parameters.Add("pre", MySqlDbType.VarChar).Value = "EAC";
            cmd.Parameters.Add("idnt", MySqlDbType.VarChar).Value = "1";
            cmd.Parameters.Add("status", MySqlDbType.Int32).Direction = ParameterDirection.Output;
            cmd.Parameters.Add("error_message", MySqlDbType.VarChar).Direction = ParameterDirection.Output;

            // Then, we create and assign a transaction to the command.
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
                const int expectedStatus = 1;
                Assert.AreEqual(expectedStatus, status);
                Assert.AreEqual("Can't delete outcome that has an association.", errorMessage);
            }
            finally
            {
                result.Close();
                // Cleanup:
                // Undo the delete by rolling back the transaction.
                transaction.Rollback();
            }
        }

        [TestMethod()]
        public void Delete_Invalid_InvalidPrefix()
        {
            // Assert:
            // Prepare the delete command and transaction.

            // First, we create the delete command.
            MySqlCommand cmd = new MySqlCommand();
            cmd.Connection = GetConnectionObject();
            cmd.CommandType = System.Data.CommandType.StoredProcedure;
            cmd.CommandText = "outcomes__delete_outcome__sp";
            cmd.Parameters.Add("pre", MySqlDbType.VarChar).Value = "Bad";
            cmd.Parameters.Add("idnt", MySqlDbType.VarChar).Value = "1";
            cmd.Parameters.Add("status", MySqlDbType.Int32).Direction = ParameterDirection.Output;
            cmd.Parameters.Add("error_message", MySqlDbType.VarChar).Direction = ParameterDirection.Output;

            // Then, we create and assign a transaction to the command.
            MySqlTransaction transaction = cmd.Connection.BeginTransaction();
            cmd.Transaction = transaction;

            // Act:
            // Execute the query.
            MySqlDataReader result;
            try 
            {
                result = cmd.ExecuteReader();
                result.Close();
            }
            catch (MySqlException e)
            {
                Trace.WriteLine(e.GetType());
            }
            finally
            {
                // Cleanup:
                // Undo the delete by rolling back the transaction.
                transaction.Rollback();
            }
        }

        [TestMethod()]
        public void Delete_Invalid_NoOutcomeWithIdentifier()
        {
            // Assert:
            // Prepare the delete command and transaction.

            // First, we create the delete command.
            MySqlCommand cmd = new MySqlCommand();
            cmd.Connection = GetConnectionObject();
            cmd.CommandType = System.Data.CommandType.StoredProcedure;
            cmd.CommandText = "outcomes__delete_outcome__sp";
            cmd.Parameters.Add("pre", MySqlDbType.VarChar).Value = "CAC";
            cmd.Parameters.Add("idnt", MySqlDbType.VarChar).Value = "9999";
            cmd.Parameters.Add("status", MySqlDbType.Int32).Direction = ParameterDirection.Output;
            cmd.Parameters.Add("error_message", MySqlDbType.VarChar).Direction = ParameterDirection.Output;

            // Then, we create and assign a transaction to the command.
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
                const int expectedStatus = 1;
                Assert.AreEqual(expectedStatus, status);
                Assert.AreEqual("Outcome does not exist.", errorMessage);
            }
            finally
            {
                result.Close();
                // Cleanup:
                // Undo the delete by rolling back the transaction.
                transaction.Rollback();
            }
        }
    }

}