import pandas as pd
import numpy as np
from matplotlib import pyplot as plt
import time
from statsmodels.tsa.arima.model import ARIMA
from sklearn.metrics import mean_squared_error
import cr_scaling

data_path = "./data.csv"

update_interval = 300 # 5MIN
min_val = 1100
max_val = 1600
max_replicas = 15

print ('Analyzer Started. (Forecast every', update_interval, 'seconds)')
def predict_new_data():
    df_predict = pd.read_csv(data_path)
    df_predict['Timestamp'] = pd.to_datetime(df_predict['Timestamp'])
    df_predict['Memory'] = df_predict['Memory'].str.split("Mi").apply(lambda x: int(x[0]))
    grouped = df_predict.groupby('Timestamp')
    node_stats = grouped.first().reset_index(drop=True)
    df_predict = node_stats
    X_predict = df_predict['Memory'].values

    model = ARIMA(X_predict, order=(5,1,0))
    model_fit = model.fit()
    
    prediction = model_fit.forecast()
    print ('\nNext Prediction Value:', prediction)

    bins = np.linspace(min_val, max_val, max_replicas)
    replica_count = np.digitize(prediction, bins, right=True)
    replica_count = np.clip(replica_count, 1, max_replicas)
    print ('Required Replica Count:', replica_count)

    cr_scaling.create_scaling_resource('test-app', 'default', int(replica_count[0]))

while True:
    predict_new_data()
    time.sleep(update_interval)
