using System;
using System.Data;
using AbOut_Database_Testing;
using Microsoft.VisualStudio.TestTools.UnitTesting;
using MySql.Data.MySqlClient;

namespace Outcomes_Testing
{
    [TestClass()]
    public class UpdateOutcome : TestConnectorMySQL
    {
        [TestMethod]
        public void Update_GoodInput()
        {
            // Arrange:
            // Prepare the query for fetching a single program.
            MySqlCommand cmd = new MySqlCommand();
            cmd.Connection = GetConnectionObject();
            cmd.CommandType = System.Data.CommandType.StoredProcedure;
            cmd.CommandText = "outcomes__update_outcome__sp";
            cmd.Parameters.Add("o_prefix", MySqlDbType.VarChar).Value = "EAC";
            cmd.Parameters.Add("o_identifier", MySqlDbType.VarChar).Value = "1";
            cmd.Parameters.Add("new_text", MySqlDbType.VarChar).Value = "Some witty phrase written at 10:40 on mini spring break.";
            cmd.Parameters.Add("status", MySqlDbType.Int32).Direction = ParameterDirection.Output;
            cmd.Parameters.Add("error_message", MySqlDbType.VarChar).Direction = ParameterDirection.Output;

            MySqlTransaction transaction = cmd.Connection.BeginTransaction();
            cmd.Transaction = transaction;

            // Act:
            // Execute the query.
            try
            {
                cmd.ExecuteNonQuery();

                // Assert:
                // Compare expected with returned results.
                int status = Convert.ToInt32(cmd.Parameters["status"].Value);
                string errorMessage = Convert.ToString(cmd.Parameters["error_message"].Value);

                const int expectedStatus = 0;
                Assert.AreEqual(expectedStatus, status);

                Assert.AreEqual("", errorMessage);

                const string tquery = "SELECT text FROM outcomes " +
                    "WHERE prefix_id = 1 AND identifier = 1";
                cmd.Connection = GetConnectionObject();
                cmd.CommandText = tquery;
                cmd.CommandType = System.Data.CommandType.Text;

                string updatedText = Convert.ToString(cmd.ExecuteScalar());

                Assert.AreEqual("Some witty phrase written at 10:40 on mini spring break.",
                    updatedText);
            }
            finally
            {
                transaction.Rollback();
            }
        }

        [TestMethod]
        public void Update_InvalidInput_PrefixNotFound()
        {
            // Arrange:
            // Prepare the query for fetching a single program.
            MySqlCommand cmd = new MySqlCommand();
            cmd.Connection = GetConnectionObject();
            cmd.CommandType = System.Data.CommandType.StoredProcedure;
            cmd.CommandText = "outcomes__update_outcome__sp";
            cmd.Parameters.Add("o_prefix", MySqlDbType.VarChar).Value = "DAK:)";
            cmd.Parameters.Add("o_identifier", MySqlDbType.VarChar).Value = "1";
            cmd.Parameters.Add("new_text", MySqlDbType.VarChar).Value = "Less clever phrase.";
            cmd.Parameters.Add("status", MySqlDbType.Int32).Direction = ParameterDirection.Output;
            cmd.Parameters.Add("error_message", MySqlDbType.VarChar).Direction = ParameterDirection.Output;

            MySqlTransaction transaction = cmd.Connection.BeginTransaction();
            cmd.Transaction = transaction;

            // Act:
            // Execute the query.
            try
            {
                cmd.ExecuteNonQuery();

                // Assert:
                // Compare expected with returned results.
                int status = Convert.ToInt32(cmd.Parameters["status"].Value);
                string errorMessage = Convert.ToString(cmd.Parameters["error_message"].Value);

                const int expectedStatus = 1;
                Assert.AreEqual(expectedStatus, status);

                Assert.AreEqual("unknown outcome prefix and/or outcome identifier", errorMessage);
            }
            finally
            {
                transaction.Rollback();
            }
        }

        [TestMethod]
        public void Update_InvalidInput_IdentifierNotFound()
        {
            // Arrange:
            // Prepare the query for fetching a single program.
            MySqlCommand cmd = new MySqlCommand();
            cmd.Connection = GetConnectionObject();
            cmd.CommandType = System.Data.CommandType.StoredProcedure;
            cmd.CommandText = "outcomes__update_outcome__sp";
            cmd.Parameters.Add("o_prefix", MySqlDbType.VarChar).Value = "EAC";
            cmd.Parameters.Add("o_identifier", MySqlDbType.VarChar).Value = "720";
            cmd.Parameters.Add("new_text", MySqlDbType.VarChar).Value = "Less clever phrase.";
            cmd.Parameters.Add("status", MySqlDbType.Int32).Direction = ParameterDirection.Output;
            cmd.Parameters.Add("error_message", MySqlDbType.VarChar).Direction = ParameterDirection.Output;

            MySqlTransaction transaction = cmd.Connection.BeginTransaction();
            cmd.Transaction = transaction;

            // Act:
            // Execute the query.
            try
            {
                cmd.ExecuteNonQuery();

                // Assert:
                // Compare expected with returned results.
                int status = Convert.ToInt32(cmd.Parameters["status"].Value);
                string errorMessage = Convert.ToString(cmd.Parameters["error_message"].Value);

                const int expectedStatus = 1;
                Assert.AreEqual(expectedStatus, status);

                Assert.AreEqual("unknown outcome prefix and/or outcome identifier", errorMessage);
            }
            finally
            {
                transaction.Rollback();
            }
        }

        [TestMethod]
        public void Update_InvalidInput_NullNewText()
        {
            // Arrange:
            // Prepare the query for fetching a single program.
            MySqlCommand cmd = new MySqlCommand();
            cmd.Connection = GetConnectionObject();
            cmd.CommandType = System.Data.CommandType.StoredProcedure;
            cmd.CommandText = "outcomes__update_outcome__sp";
            cmd.Parameters.Add("o_prefix", MySqlDbType.VarChar).Value = "EAC";
            cmd.Parameters.Add("o_identifier", MySqlDbType.VarChar).Value = "1";
            cmd.Parameters.Add("new_text", MySqlDbType.VarChar).Value = null;
            cmd.Parameters.Add("status", MySqlDbType.Int32).Direction = ParameterDirection.Output;
            cmd.Parameters.Add("error_message", MySqlDbType.VarChar).Direction = ParameterDirection.Output;

            MySqlTransaction transaction = cmd.Connection.BeginTransaction();
            cmd.Transaction = transaction;

            // Act:
            // Execute the query.
            try
            {
                cmd.ExecuteNonQuery();

                // Assert:
                // Compare expected with returned results.
                int status = Convert.ToInt32(cmd.Parameters["status"].Value);
                string errorMessage = Convert.ToString(cmd.Parameters["error_message"].Value);

                const int expectedStatus = 1;
                Assert.AreEqual(expectedStatus, status);

                Assert.AreEqual("encountered null new text", errorMessage);
            }
            finally
            {
                transaction.Rollback();
            }
        }

        [TestMethod]
        public void Update_InvalidInput_NullIdentifier()
        {
            // Arrange:
            // Prepare the query for fetching a single program.
            MySqlCommand cmd = new MySqlCommand();
            cmd.Connection = GetConnectionObject();
            cmd.CommandType = System.Data.CommandType.StoredProcedure;
            cmd.CommandText = "outcomes__update_outcome__sp";
            cmd.Parameters.Add("o_prefix", MySqlDbType.VarChar).Value = "EAC";
            cmd.Parameters.Add("o_identifier", MySqlDbType.VarChar).Value = null;
            cmd.Parameters.Add("new_text", MySqlDbType.VarChar).Value = "No effort here.";
            cmd.Parameters.Add("status", MySqlDbType.Int32).Direction = ParameterDirection.Output;
            cmd.Parameters.Add("error_message", MySqlDbType.VarChar).Direction = ParameterDirection.Output;

            MySqlTransaction transaction = cmd.Connection.BeginTransaction();
            cmd.Transaction = transaction;

            // Act:
            // Execute the query.
            try
            {
                cmd.ExecuteNonQuery();

                // Assert:
                // Compare expected with returned results.
                int status = Convert.ToInt32(cmd.Parameters["status"].Value);
                string errorMessage = Convert.ToString(cmd.Parameters["error_message"].Value);

                const int expectedStatus = 1;
                Assert.AreEqual(expectedStatus, status);

                Assert.AreEqual("unknown outcome prefix and/or outcome identifier", errorMessage);
            }
            finally
            {
                transaction.Rollback();
            }
        }

        [TestMethod]
        public void Update_InvalidInput_NullPrefix()
        {
            // Arrange:
            // Prepare the query for fetching a single program.
            MySqlCommand cmd = new MySqlCommand();
            cmd.Connection = GetConnectionObject();
            cmd.CommandType = System.Data.CommandType.StoredProcedure;
            cmd.CommandText = "outcomes__update_outcome__sp";
            cmd.Parameters.Add("o_prefix", MySqlDbType.VarChar).Value = null;
            cmd.Parameters.Add("o_identifier", MySqlDbType.VarChar).Value = "1";
            cmd.Parameters.Add("new_text", MySqlDbType.VarChar).Value = "No effort here.";
            cmd.Parameters.Add("status", MySqlDbType.Int32).Direction = ParameterDirection.Output;
            cmd.Parameters.Add("error_message", MySqlDbType.VarChar).Direction = ParameterDirection.Output;

            MySqlTransaction transaction = cmd.Connection.BeginTransaction();
            cmd.Transaction = transaction;

            // Act:
            // Execute the query.
            try
            {
                cmd.ExecuteNonQuery();

                // Assert:
                // Compare expected with returned results.
                int status = Convert.ToInt32(cmd.Parameters["status"].Value);
                string errorMessage = Convert.ToString(cmd.Parameters["error_message"].Value);

                const int expectedStatus = 1;
                Assert.AreEqual(expectedStatus, status);

                Assert.AreEqual("unknown outcome prefix and/or outcome identifier", errorMessage);
            }
            finally
            {
                transaction.Rollback();
            }
        }

    }

}