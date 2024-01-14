Running the following commands from the base directory (directory where the readme resides) will verify the respective generated encoding using the provided Gobra jar.


java -Xss128m -jar /Users/arquintlinard/ETH/PhD/gobra/target/scala-2.13/gobra.jar -I ./ --module github.com/aws/amazon-ssm-agent/agent/iospecs -i place/place.gobra
java -Xss128m -jar /Users/arquintlinard/ETH/PhD/gobra/target/scala-2.13/gobra.jar -I ./ --module github.com/aws/amazon-ssm-agent/agent/iospecs -i fresh/fresh.gobra
java -Xss128m -jar /Users/arquintlinard/ETH/PhD/gobra/target/scala-2.13/gobra.jar -I ./ --module github.com/aws/amazon-ssm-agent/agent/iospecs -i pub/pub.gobra
java -Xss128m -jar /Users/arquintlinard/ETH/PhD/gobra/target/scala-2.13/gobra.jar -I ./ --module github.com/aws/amazon-ssm-agent/agent/iospecs -i term/term.gobra
java -Xss128m -jar /Users/arquintlinard/ETH/PhD/gobra/target/scala-2.13/gobra.jar -I ./ --module github.com/aws/amazon-ssm-agent/agent/iospecs -i bytes/bytes.gobra
java -Xss128m -jar /Users/arquintlinard/ETH/PhD/gobra/target/scala-2.13/gobra.jar -I ./ --module github.com/aws/amazon-ssm-agent/agent/iospecs -i claim/claim.gobra
java -Xss128m -jar /Users/arquintlinard/ETH/PhD/gobra/target/scala-2.13/gobra.jar -I ./ --module github.com/aws/amazon-ssm-agent/agent/iospecs -i fact/fact.gobra
java -Xss128m -jar /Users/arquintlinard/ETH/PhD/gobra/target/scala-2.13/gobra.jar -I ./ --module github.com/aws/amazon-ssm-agent/agent/iospecs -i iospec/permissions_out.gobra
java -Xss128m -jar /Users/arquintlinard/ETH/PhD/gobra/target/scala-2.13/gobra.jar -I ./ --module github.com/aws/amazon-ssm-agent/agent/iospecs -i iospec/permissions_in.gobra
java -Xss128m -jar /Users/arquintlinard/ETH/PhD/gobra/target/scala-2.13/gobra.jar -I ./ --module github.com/aws/amazon-ssm-agent/agent/iospecs -i iospec/permissions_Agent_internal.gobra
java -Xss128m -jar /Users/arquintlinard/ETH/PhD/gobra/target/scala-2.13/gobra.jar -I ./ --module github.com/aws/amazon-ssm-agent/agent/iospecs -i iospec/permissions_in.gobra iospec/permissions_out.gobra iospec/permissions_Agent_internal.gobra iospec/Agent.gobra