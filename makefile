.PHONY: run-images testnet mysql minibank

run-images: testnet mysql minibank  


minibank: 
	docker build -t minibank ./minibank
	docker run -i -t -d --name minibank -p 8080:8080 --net testnet minibank

mysql:
	docker build -t mysql ./mysql
	docker run -i -t -d --name mysql --net testnet mysql	
	sleep 10

testnet:
	docker network create testnet

clean:
	docker stop minibank
	docker rm minibank
	docker stop mysql
	docker rm mysql
	docker network rm testnet
	
remove-unused-ICVN: 
	docker system prune -a

