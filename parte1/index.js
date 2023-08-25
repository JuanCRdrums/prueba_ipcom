const express = require('express');
const axios = require('axios');

const app = express();
const port = 3000;
const service_url = "https://apirecruit-gjvkhl2c6a-uc.a.run.app/compras/";

app.get('/resumen/:date', async (req, res) => {
    try {
        //data fetching
        let all_data = [];
        for(let i = 0; i < req.query.dias; i++)
        {
            const date = new Date(req.params.date);
            date.setDate(date.getDate() + i);
            const string_date = date.toISOString().split('T')[0];
            const response = await axios.get(service_url + string_date);
            response.data.forEach(element => {
                all_data.push(element);
            });
        }

        //data processing
        let total = 0;
        let nocompraron = 0;
        let compraMasAlta = 0;
        all_data.forEach(element => {
            if(element.monto !== undefined) total += element.monto;
            if(!element.compro) nocompraron++;
            if(element.monto > compraMasAlta) compraMasAlta = element.monto;
        });

        let comprasPorTDCMap = all_data.reduce((map, element) => {
            const tdc = element.tdc;
            if(tdc !== undefined) map.set(tdc, (map.get(tdc) || 0) + 1);
            return map;
          }, new Map());
        
        const comprasPorTDC = {};
        for (const [key, value] of comprasPorTDCMap) {
            comprasPorTDC[key] = value;
        }


        let ret_value = {
            "total": total,
            "comprasPorTDC": comprasPorTDC,
            "nocompraron": nocompraron,
            "compraMasAlta": compraMasAlta
        };

        
        
        res.json(ret_value);
    } catch (error) {
        console.error('Error getting the data:', error);
        res.status(500).json({ error: 'Internal server error' });
    }
  });

app.listen(port, () => {
    console.log(`Server running in port ${port}`);
});
  
  