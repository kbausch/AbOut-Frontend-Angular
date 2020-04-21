using Microsoft.VisualStudio.TestTools.UnitTesting;
using MySql.Data.MySqlClient;
using System;

namespace AbOut_Database_Testing
{
    public class TestConnectorMySQL
    {
        // This needs to be pulled from a configuration file.
        private const string connectionString = "server=127.0.0.1;port=3306;database=assessment;uid=web_user;pwd=resu_bew";
        private static MySqlConnection connectionSingleton = null;
        
        protected static MySqlConnection GetConnectionObject()
        {
            if (connectionSingleton == null)
            {
                connectionSingleton = new MySqlConnection(connectionString);
                connectionSingleton.Open();
            }
            return connectionSingleton;
        }

        [TestCleanup()]
        public void TestCleanUp()
        {
            connectionSingleton.Close();
            connectionSingleton = null;
        }
    }
}