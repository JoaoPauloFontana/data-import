import React, { useState, useEffect } from 'react';

const Dashboard = ({ token }) => {
    const [records, setRecords] = useState([]);
    const [summary, setSummary] = useState({});
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);

    useEffect(() => {
        const fetchData = async () => {
            setLoading(true);
            setError(null);
            try {
                const response = await fetch('http://localhost:8200/api/records', {
                    headers: {
                        Authorization: `Bearer ${token}`,
                    },
                });

                if (!response.ok) {
                    throw new Error(`HTTP error! status: ${response.status}`);
                }

                const data = await response.json();
                setRecords(data);
                calculateSummary(data);
            } catch (error) {
                setError("Erro ao buscar registros");
                console.error("Erro ao buscar registros", error);
            } finally {
                setLoading(false);
            }
        };

        fetchData();
    }, [token]);

    const calculateSummary = (data) => {
        const summary = data.reduce((acc, record) => {
            acc.total += record.billing_pre_tax_total;
            return acc;
        }, { total: 0 });

        setSummary(summary);
    };

    if (loading) {
        return <div className="loading">Carregando...</div>;
    }

    if (error) {
        return <div className="error">{error}</div>;
    }

    return (
        <div className="app-container">
            <h1>Dashboard</h1>
            <div>Total Billing Pre Tax: {summary.total}</div>
            <div className="table-container">
                <table>
                    <thead>
                        <tr>
                            <th>Partner ID</th>
                            <th>Customer Name</th>
                            <th>Billing Pre Tax Total</th>
                            <th>Month</th>
                        </tr>
                    </thead>
                    <tbody>
                        {records.map(record => (
                            <tr key={record.id}>
                                <td>{record.partner_id}</td>
                                <td>{record.customer_name}</td>
                                <td>{record.billing_pre_tax_total}</td>
                                <td>{new Date(record.usage_date).toLocaleDateString('en-US', { month: 'long' })}</td>
                            </tr>
                        ))}
                    </tbody>
                </table>
            </div>
        </div>
    );
};

export default Dashboard;
