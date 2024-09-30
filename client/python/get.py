import sys
import grpc

sys.path.append('../../../microservices-proto/python/order')

import order_pb2
import order_pb2_grpc

def run():
    with grpc.insecure_channel('127.0.0.1:3000') as channel:
        client = order_pb2_grpc.OrderStub(channel)

        request = order_pb2.GetOrderRequest(
            order_id=3
        )

        response = client.Get(request)

        print("Order created:", response)

if __name__ == '__main__':
    run()
